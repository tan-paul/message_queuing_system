package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	database "message_queuing_system/database"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	pngquant "github.com/yusukebe/go-pngquant"
)

func consumer() {
	ch, err := database.RabbitmqConn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	product_id, ok, err := ch.Get("Product_Message_Queue", false)
	if ok {
		ch.Ack(product_id.DeliveryTag, false)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Product id not found")
		return
	}
	pid := 0
	err = json.Unmarshal(product_id.Body, &pid)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product id recivied from queue :: ", pid)
	res, err := database.Database_Conn.Query("select product_images from Image_Store where product_id=?", strconv.Itoa(pid))
	if err != nil {
		fmt.Println("Error while reading data from DB")
		return
	}
	image_url, index := "", 0
	for res.Next() {
		err := res.Scan(&image_url)
		if err != nil {
			fmt.Println("Error while reading data from DB")
			return
		}
		local_path, err := downloadFile(image_url, "image_"+strconv.Itoa(index)+"_"+strconv.Itoa(pid))
		if err != nil {
			fmt.Println("Error while downloading Image with url :: ", image_url)
			fmt.Println("error :: ", err)
			return
		}
		_, err = database.Database_Conn.Exec("update Image_Store set compressed_product_images=? where  product_id=? and product_images=?", local_path, pid, image_url)
		if err != nil {
			fmt.Println("Error while updating db")
			return
		}
		_, err = database.Database_Conn.Exec("update Products set updated_at=NOW() where  product_id=?", pid)
		if err != nil {
			fmt.Println("Error while updating db")
			return
		}
		index += 1
		image_url = ""
	}
}

func downloadFile(URL, fileName string) (string, error) {
	path := "/Users/tanmoy.p/Downloads/images/" + fileName

	_, err := url.ParseRequestURI(URL)
	if err != nil {
		URL = "https://" + URL
		_, err = url.ParseRequestURI(URL)
		if err != nil {
			return "invalid url", err
		}
	}
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New("Received non 200 response code")
	}
	if strings.Contains(URL, ".jpg") || strings.Contains(URL, ".jpeg") {
		return optimizeJPG(response, path)
	} else if strings.Contains(URL, ".png") {
		return optimizePNG(response, path)
	} else {
		return "", errors.New("invalid or unsupported format")
	}
}

func optimizeJPG(response *http.Response, path string) (string, error) {
	file, err := os.Create(path + ".jpg")
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	in := bytes.NewReader(data)

	img, err := jpeg.Decode(in)
	if err != nil {
		return "", err
	}

	// Encode image
	out := new(bytes.Buffer)

	err = jpeg.Encode(out, img, &jpeg.Options{
		Quality: 20,
	})
	if err != nil {
		return "", err
	}

	outlen := int64(out.Len())
	_, err = io.Copy(file, out)
	if err != nil {
		return "", err
	}
	if outlen < in.Size() {
		saved := (in.Size() - outlen) * 100 / in.Size()
		fmt.Println("imaged compressed " + fmt.Sprintf("%02d%%", saved))
	}
	return path + ".jpg", nil
}
func optimizePNG(response *http.Response, path string) (string, error) {
	file, err := os.Create(path + ".png")
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	in := bytes.NewReader(data)

	img, err := png.Decode(in)
	if err != nil {
		return "", err
	}

	// Encode image
	out := new(bytes.Buffer)
	cimg, err := pngquant.Compress(img, "1")
	if err != nil {
		return "", err
	}
	err = png.Encode(out, cimg)
	if err != nil {
		return "", err
	}

	outlen := int64(out.Len())
	_, err = io.Copy(file, out)
	if err != nil {
		return "", err
	}

	if outlen < in.Size() {
		saved := (in.Size() - outlen) * 100 / in.Size()
		fmt.Println("imaged compressed " + fmt.Sprintf("%02d%%", saved))
	}
	return path + ".png", nil
}

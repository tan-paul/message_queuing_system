package main

import (
	"testing"
)

type downloadFileTest struct {
	url, fileName, expected, err string
}

var downloadFileTests = []downloadFileTest{
	// test case for .jpg or .jpeg extension
	{"https://images.pexels.com/photos/247502/pexels-photo-247502.jpeg?cs=srgb&dl=pexels-pixabay-247502.jpg&fm=jpg&_gl=1*sclgt5*_ga*OTc3MTI2MjIxLjE2OTM4OTc2MDg.*_ga_8JE65Q40S6*MTY5NDA2NjgzNy4zLjEuMTY5NDA2Njg0Mi4wLjAuMA..", "jpg_file", "/Users/tanmoy.p/Downloads/images/jpg_file.jpg", ""},
	{"https://cdn.pixabay.com/photo/2012/06/19/10/32/owl-50267_1280.jpg", "jpg_file", "/Users/tanmoy.p/Downloads/images/jpg_file.jpg", ""},
	{"https://images.pexels.com/photos/68421/pexels-photo-68421.jpeg?cs=srgb&dl=pexels-piet-bakker-68421.jpg&fm=jpg&_gl=1*h0wu9z*_ga*OTc3MTI2MjIxLjE2OTM4OTc2MDg.*_ga_8JE65Q40S6*MTY5NDI2NzE2Mi40LjEuMTY5NDI2NzE2NS4wLjAuMA..", "jpg_file", "/Users/tanmoy.p/Downloads/images/jpg_file.jpg", ""},

	//test case for .png extension
	{"https://image.similarpng.com/very-thumbnail/2020/09/Raw-beef-meat-pieces-on-transparent-background-PNG.png", "png_file", "/Users/tanmoy.p/Downloads/images/png_file.png", ""},
	{"https://www.techsmith.com/blog/wp-content/uploads/2020/11/TechSmith-Blog-JPGvsPNG.png", "png_file", "/Users/tanmoy.p/Downloads/images/png_file.png", ""},
	{"https://image.similarpng.com/very-thumbnail/2020/05/Cartoon-chef-holding-hamburger-transparent-background-PNG.png", "png_file", "/Users/tanmoy.p/Downloads/images/png_file.png", ""},
	{"http://image.similarpng.com/very-thumbnail/2020/05/Cartoon-chef-holding-hamburger-transparent-background-PNG.png", "png_file", "/Users/tanmoy.p/Downloads/images/png_file.png", ""},
	{"image.similarpng.com/very-thumbnail/2020/05/Cartoon-chef-holding-hamburger-transparent-background-PNG.png", "png_file", "/Users/tanmoy.p/Downloads/images/png_file.png", ""},

	// test case for other types
	{"https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf", "test_file", "", "invalid or unsupported format"},
	{"www.google.com", "test_file", "", "invalid or unsupported format"},

	//invalid url test
	{"www.goo.gle.com", "invalid_file", "", "Get \"https://www.goo.gle.com\": tls: failed to verify certificate: x509: “sni-support-required-for-valid-ssl” certificate is not standards compliant"},
	//{"ttt@fdsf", "invalid_file", "invalid url", "Get \"https://ttt@fdsf\": dial tcp: lookup fdsf: no such host"},
}

func TestDownloadFile(t *testing.T) {
	for _, test := range downloadFileTests {
		output, err := downloadFile(test.url, test.fileName)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
		if err != nil && err.Error() != test.err {
			t.Errorf("error %q not equal to expected %q", err, test.err)
		}

	}
}

func BenchmarkDownloadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		downloadFile("https://www.techsmith.com/blog/wp-content/uploads/2020/11/TechSmith-Blog-JPGvsPNG.png", "benchmark_file")
	}
}

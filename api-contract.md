# api contract for Message Queuing System

## url :- 
```
POST http://localhost:8080/product_info
```

## sample I/p I

```json
{
    "user_id":2,
    "product_name":"chukchuki",
    "product_description":"it's afunny product for babies",
    "product_images":[
        "https://www.techsmith.com/blog/wp-content/uploads/2020/11/TechSmith-Blog-JPGvsPNG.png",
        "https://image.similarpng.com/very-thumbnail/2020/09/Raw-beef-meat-pieces-on-transparent-background-PNG.png",
        "https://image.similarpng.com/very-thumbnail/2020/05/Floating-burger-transparent-background-PNG.png",
        "https://image.similarpng.com/very-thumbnail/2020/05/Cartoon-chef-holding-hamburger-transparent-background-PNG.png",
        "https://cdn.pixabay.com/photo/2012/06/19/10/32/owl-50267_1280.jpg",
        "https://images.pexels.com/photos/68421/pexels-photo-68421.jpeg?cs=srgb&dl=pexels-piet-bakker-68421.jpg&fm=jpg&_gl=1*h0wu9z*_ga*OTc3MTI2MjIxLjE2OTM4OTc2MDg.*_ga_8JE65Q40S6*MTY5NDI2NzE2Mi40LjEuMTY5NDI2NzE2NS4wLjAuMA..",
        "https://images.pexels.com/photos/1152077/pexels-photo-1152077.jpeg?cs=srgb&dl=pexels-ge-yonk-1152077.jpg&fm=jpg&_gl=1*wtfcw0*_ga*OTc3MTI2MjIxLjE2OTM4OTc2MDg.*_ga_8JE65Q40S6*MTY5Mzg5NzYxMC4xLjEuMTY5Mzg5NzYyMi4wLjAuMA..",
        "https://images.pexels.com/photos/247502/pexels-photo-247502.jpeg?cs=srgb&dl=pexels-pixabay-247502.jpg&fm=jpg&_gl=1*sclgt5*_ga*OTc3MTI2MjIxLjE2OTM4OTc2MDg.*_ga_8JE65Q40S6*MTY5NDA2NjgzNy4zLjEuMTY5NDA2Njg0Mi4wLjAuMA..",
        "https://image.similarpng.com/thumbnail/2022/07/Potato-chips-in-a-glass-plate-on-transparent-background-PNG.png"
    ],
    "product_price":3456
}
```
## sample o/p I

### status code :- 201Created

```json
{
    "Message": "Queued"
}
```

## sample I/p II

```json
{
    "user_id":1,
    "product_name":"chukchuki",
    "product_description":"it's afunny product for babies",
    "product_price":3765678
}
```
## sample o/p II

### status :- 400Bad Request

```json
{
    "Error": "Product image urls required"
}
```
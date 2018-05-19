# img_gen:使用Golang自定义文字生成表情包图片.

## Installation
    go get github.com/wenjiax/img_gen
    go build

## Example
    ./img_gen -s "今天在群里看到老是有人发这样的图片，我看着心里很不爽，于是写个程序来装逼。现在终于写完了，我要去装逼了，这个逼我记住了..."
    ![](https://raw.githubusercontent.com/wenjiax/img_gen/master/out.png)

    ./img_gen -o out1.png -t temp2.png -s "今天在群里看到老是有人发这样的图片，我看着心里很不爽，于是写个程序来装逼。现在终于写完了，我要去装逼了，这个逼我记住了..."
    ![](https://raw.githubusercontent.com/wenjiax/img_gen/master/out2.png)
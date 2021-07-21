# S3 Pre-Signed URL AWS Lambda

AWS Lambda Function that generates S3 Pre-Signed URL for the specified item

### Function Payload

```json
{
  "bucket-name": "enter the bucket's name here",
  "item-name": "enter the item's name here"
}
```

### Function Output

```json
{
  "presigned-url": "the presigned url of the requested item"
}
```

## Usage

1. Clone this repository

```sh
git clone https://github.com/pholawat-tle/S3-PreSignedURL-Lambda
```

2. Run the command below to build the binary

```sh
GOOS=linux GOARCH=amd64 go build main.go
```

3. Run the command below to compress the binary into a .zip file

```sh
zip function.zip main
```

4. Create a Lambda Function and attached an IAM Role that allow s3:GetObject API call

5. _If necessary,_ edit the S3 Bucket Policy to allow access from Lambda IAM Role

6. Upload function.zip to AWS Lambda

_This function is configured for S3 buckets in ap-southeast-1 region_

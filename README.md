# stori

## Overview

This tool is designed for reading, validating, and reporting client transaction data from CSV files. Each transaction record in the CSV file encapsulates information about credit and debit movements. After processing, the tool validates each entry, generates a detailed report of client movements, sends the report via email (with the original file attached), and stores the information in a database.

## Features

* CSV File Processing: Efficiently reads and processes CSV files containing client transaction information.

* Entry Validation: Implements robust validation checks to ensure data integrity and accuracy.

* Report Generation: Creates a comprehensive report of client movements, summarizing credit and debit details.

* Email Notification: Sends an email with the generated report attached, providing a convenient way to review the processed data.

* Database Storage: Persists client movement information in a specified database for future reference.

* Multiple User Interfaces:
    * CLI Tool: Enables users to interact with the tool via the command line interface.
    * REST API (Dockerized): Exposes a REST API using Docker for seamless integration into existing systems.
    * AWS Lambda Function: Provides a serverless option for processing transactions in an AWS environment.


## Environment Variables

```
STORI_EMAIL
STORI_POSTGRES_DBNAME
STORI_POSTGRES_HOST
STORI_POSTGRES_PASSWORD
STORI_POSTGRES_PORT
STORI_POSTGRES_USER
STORI_SENDER_API_KEY
```

## Email setup

You need to setup an API Key for [resend](https://resend.com/).

## Usage

1. Set up Environment Variables.

2. Start the Container with Docker Compose.

```shell
$ docker compose up
```

#### CLI Tool

Compile the binary.

```shell
$ go build -o stori cmd/cli/cli.go
```
Execute the tool.

```shell
$ ./stori -file input.csv -email your@email.com
```

#### REST API

Make a GET request to your `localhost:3000` passing your email as a URL param.

```shell
$ curl -XGET 'http://localhost:3000/your@email.com'
```

#### AWS Lambda

You can use the [GitHub Actions](https://docs.github.com/en/actions) and [Terraform](https://www.terraform.io/) config files to deploy the application.

The lambda should react to a S3 bucket. So when a new csv file is uploaded the tool generate the report and send the email.

## Tests

You can run the tests executing the following command:

```shell
$ make coverage
```

This will give you information about the tests coverage on your terminal and your web browser.

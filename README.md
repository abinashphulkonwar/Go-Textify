# Textify

This tool utilizes computer vision to extract text from images.

## Set Up

### Prerequisites

Before you begin, ensure you have the following:

1. An Azure account with an active subscription.
2. An Azure Computer Vision resource set up.
3. Obtain the API Key and URL of your Azure Computer Vision resource, and store them in a `config.json` file as shown below:

```json
{
  "api_key": "key",
  "api_url": "url"
}
```

4. Copy the absolute path of the `config.json` file and set it in the `app/handlers/extract.go`` file at line number 25, like this:

```go
data, err := os.ReadFile("path_to_config") // Replace "path_to_config" with the absolute path to your `config.json` file as mentioned in point 3.
```

## Build

To build Textify, run the following command:

```sh
go build
```

## Setting the Path

You can add the tool's path to your OS environment variables to access it from anywhere.

## Usage

To extract text from an image, use the following command:

```sh
Textify -i="./20230816_133440.jpg" -o="output.text"
```

For Help:

```sh
Textify help
```

## Contribution

Improvements made:

1. Clarified the purpose of the tool in the project description.
2. Improved the formatting of the setup steps.
3. Clarified the usage of the `Textify` command.
4. Provided a more consistent naming convention for the output file.
5. Added a "Contribution" section, which appears to be missing in your original README.
6. Enhanced the overall readability and organization of the document.

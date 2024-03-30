<h1 align="center">Analyze-Tags</h1>

<p align="center">
  <a href="https://pkg.go.dev/github.com/mtnmunuklu/analyze-tags">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-informational.svg" alt="Go Doc">
  </a>
  <a href="https://goreportcard.com/report/github.com/mtnmunuklu/analyze-tags">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A+-success.svg" alt="Go Report">
  </a>
  <!-- Other links and badges -->
</p>


Analyze-Tags is a tool designed for analyzing security rules and generating charts and Excel files based on the provided rules. It supports various types of security rules such as Sigma, YARA, and Csiem rules.


## Table of Contents

- [Installation](#installation)
  - [Normal Installation](#normal-installation)
  - [Docker Installation](#docker-installation)
- [Usage](#usage)
  - [Command-line Flags](#command-line-flags)
  - [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use Analyze-Tags, you can choose between two installation options:

### Normal Installation

Analyze-Tags provides precompiled ZIP files for different platforms. Download the appropriate ZIP file for your platform from the following links:

- [Windows](https://github.com/mtnmunuklu/analyze-tags/releases/latest/download/analyze-tags-windows-latest.zip)
- [Linux](https://github.com/mtnmunuklu/analyze-tags/releases/latest/download/analyze-tags-ubuntu-latest.zip)
- [macOS](https://github.com/mtnmunuklu/analyze-tags/releases/latest/download/analyze-tags-macos-latest.zip)

Once downloaded, extract the ZIP file to a directory of your choice. Among the extracted files, you will find the analyze-tags executable.

Ensure that the directory containing the Analyze-Tags executable is added to your system's PATH environment variable, enabling you to run Analyze-Tags from any location in the command line.

Note: Analyze-Tags requires Go to be installed on your system. Download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)

### Docker Installation

Alternatively, you can use Docker to run Analyze-Tags in a containerized environment. Docker provides a convenient and consistent way to set up and use Analyze-Tags without worrying about dependencies or system configurations.

To install and set up Analyze-Tags using Docker, make sure Docker is installed on your system. If not, download and install Docker from the official website: [https://www.docker.com/get-started](https://www.docker.com/get-started)

Once Docker is installed, follow these steps:

1. **Clone the Repository**: If not done already, clone the Analyze-Tags repository to your local machine:

   ```shell
   git clone https://github.com/mtnmunuklu/analyze-tags.git
   ```

2. **Navigate to Docker Directory**: Go to the docker directory inside the cloned repository:
   
   ```shell
   cd tools/docker
   ```

3. **Build Docker Image and Start Container**: Use the setup script to build the Docker image named analyze-tags:

   ```shell
   go run setup_docker_analyze_tags.go -rules <rulesDirectory> -config <configFile> -output <outputDirectory>
   ```

    This script handles building the Docker image and starting the container for you.

That's it! You have successfully installed Analyze-Tags on your system. Proceed to the  [Usage](#usage) to learn how to use Analyze-Tags.

If you prefer to build Analyze-Tags from source, refer to the [Build Instructions](BUILD.md) for detailed steps on building and installing it on your platform.

## Usage

### Command-line Flags

Analyze-Tags provides several command-line flags for configuring its behavior:

- `-filepath`: Specifies the name or path of the file or directory to read.
- `-filecontent`: Specifies the base64-encoded content of the file or directory to read.
- `-output`: Specifies the output directory for writing files.
- `chart`: Specifies whether to generate charts.
- `-chartType`: Specifies one or more chart types to generate (comma-separated).
- `-excel`: Generates Excel files.
- `-sigma`, `-yara`, `-csiem`: Specifies the type of rules to use.

For more details on available flags, you can use the `-help` flag:
   ```shell
   analyze-tags -help
   ```

### Examples

Here are a few examples of using Analyze-Tags:

- To generate synthetic logs from a Sigma/Yara/Csiem rule file and a configuration file:

   ```shell
   analyze-tags -filepath /path/to/sigma/rule.yml -config /path/to/config.yml -chart -chartType "bar,line"
   ```
   or
   ```shell
   docker exec analyze-tags ./analyze-tags -filepath /path/to/sigma/rule.yml -config /path/to/config.yml -chart -chartType "bar,line"
   ```

- To generate synthetic logs from Sigma/Yara/Csiem rule content and configuration content:

   ```shell
   analyze-tags -filecontent base64_encoded_rule_content -configcontent base64_encoded_config_content -chart -chartType "bar,line"
   ```
   or
   ```shell
   docker exec analyze-tags ./analyze-tags -filecontent base64_encoded_rule_content -configcontent base64_encoded_config_content -chart -chartType "bar,line"
   ```

## Contributing

Contributions to Analyze-Tags are welcome and encouraged! Please read the [contribution guidelines](CONTRIBUTING.md) before making any contributions to the project.

## License

Analyze-Tags is licensed under the MIT License. See [LICENSE](LICENSE) for the full text of the license.
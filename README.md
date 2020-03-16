[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/F0xedb/pi-day-2020">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/42/Greek_lc_pi.svg/1200px-Greek_lc_pi.svg.png" alt="Logo" width="300" height="300">
  </a>

  <h3 align="center">Pi Day 2020</h3>

  <p align="center">
    A simple query to find the index where a number occures in pi
    <br />
    <a href="https://github.com/F0xedb/pi-day-2020"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/F0xedb/pi-day-2020">View Demo</a>
    ·
    <a href="https://github.com/F0xedb/pi-day-2020/issues">Report Bug</a>
    ·
    <a href="https://github.com/F0xedb/pi-day-2020/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->

## Table of Contents

- [About the Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgements](#acknowledgements)

<!-- ABOUT THE PROJECT -->

## About The Project

This is just a fun repository for celebrating `PI Day`!

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

All you need is golang!

```sh
pacman -Syu go
```

### Installation

1. Clone the pi-day-2020

```sh
git clone https://github.com/F0xedb/pi-day-2020.git
```

2. Install packages

```sh
pacman -Syu go
```

3. Optionaly download a much bigger file

```sh
# one billion
wget https://ia800409.us.archive.org/view_archive.php\?archive\=/9/items/Math_Constants/Pi.zip\&file\=Pi%20-%20Dec.txt -O pi.txt # roughly 1 GB
```

<!-- USAGE EXAMPLES -->

## Usage

To run everything simply launch the webserver

```sh
go run main.go
```

After seeing to output `Pi Digits loaded in!` the webserver has booted.

In case it takes to long to load you can do the following.

Open main.go and edit this line

```go
var speed = 100000
```

Increasing the number should increase the loading speed

The URL should be `localhost:8080`

Below are some screenshots of the web frontend

![Empty](/screenshot/Empty.png)

![Full](/screenshot/Full.png)

### API

The api is very simple to use and is accesable at `localhost:8080`

we have one endpoint called `/search/<number>`

Here is an example request

```
GET localhost:8080/search/123
```

With the following response

```json
{
  "Index": 1924,
  "Input": "123",
  "Surrounding": "1237137869"
}
```

<!-- ROADMAP -->

## Roadmap

See the [open issues](https://github.com/F0xedb/pi-day-2020/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

F0xedb - tom@odex.be

Project Link: [https://github.com/F0xedb/pi-day-2020](https://github.com/F0xedb/pi-day-2020)

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

- [F0xedb](https://github.com/F0xedb/pi-day-2020)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/F0xedb/pi-day-2020.svg?style=flat-square
[contributors-url]: https://github.com/F0xedb/pi-day-2020/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/F0xedb/pi-day-2020.svg?style=flat-square
[forks-url]: https://github.com/F0xedb/pi-day-2020/network/members
[stars-shield]: https://img.shields.io/github/stars/F0xedb/pi-day-2020.svg?style=flat-square
[stars-url]: https://github.com/F0xedb/pi-day-2020/stargazers
[issues-shield]: https://img.shields.io/github/issues/F0xedb/pi-day-2020.svg?style=flat-square
[issues-url]: https://github.com/F0xedb/pi-day-2020/issues
[license-shield]: https://img.shields.io/github/license/F0xedb/pi-day-2020.svg?style=flat-square
[license-url]: https://github.com/F0xedb/pi-day-2020/blob/master/LICENSE.txt
[product-screenshot]: https://upload.wikimedia.org/wikipedia/commons/thumb/4/42/Greek_lc_pi.svg/1200px-Greek_lc_pi.svg.png

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://cleverdinosaur.com">
    <img src="https://cleverdinosaur.com/wp-content/uploads/2020/12/cropped-smart-logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Dash!</h3>

  <p align="center">
    A DIGITAL CODE CRAFTING BEAUTIFUL EXPERIENCES.
    <br />
    <strong>UNIQUE. POWERFUL. CREATIVE.</strong>
    <br />
    <br />
    <a href="#">View Demo</a>
    ·
    <a href="https://github.com/cleverdinosaur/dash/issues">Report Bug</a>
    ·
    <a href="https://github.com/cleverdinosaur/dash/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://cleverdinosaur.com)

Dash! is a 2D barcode alternative similar to a QR code, but instead uses an icon or company logo and a "dash border" rather than a square pattern of black dots.

Similar to a QR code, Dash! can be used to take consumers to a brand’s website, but can also facilitate mobile purchases, coupon downloads, free sample requests, video views, promotional entries, Facebook Likes, Pinterest Pins, Twitter Follows, Posts and Tweets and so much more.

Currently the project works by encoding a color hex bar on the very first pixel of an image with the dashed border added purely for recognition that it is infact a dash code and can be scanned.

Dash 1.0 is primarily a POC and I hope to add the encoding directly within the dashed border in the near future. You may also suggest changes by forking this repo and creating a pull request or opening an issue. 



<!-- GETTING STARTED -->
## Getting Started
Dash is developed with Golang in the backend and a React web application on the frontend so you will need these and the associated technologies installed (npm, yarn, golang, react etc)
To get a local dev copy up and running follow these simple steps.


### Installation

1. Install the pre-requisites
2. Clone the repo
   ```sh
   git clone https://github.com/cleverdinosaur/dash.git
   ```
3. Start the Go Server
   ```sh
   go run main.go
   ```
4. Start the React Frontend
   ```sh
   cd frontend
   npm start
   ```
5. Visit http://localhost:3000 to view the running demo



<!-- USAGE EXAMPLES -->
## Usage

You can use the Dash backend as a headless API or use the bundled frontend out of the box.

The primary usage of Dash is pretty much anything you would use a QR code or barcode for so the possibilities are endless.

For v1.0 we recommend sticking to URL's but more features and native functionality will be added soon without the need for hacking around.

_For more examples, please refer to the [Documentation](https://cleverdinosaur.com/projects/dash)_



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/cleverdinosaur/dash/issues) for a list of proposed icons (and known issues).



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

Connor Gibson - [@thecleverdino](https://twitter.com/thecleverdino) - roar@cleverdinosaur.com

Project Link: [https://github.com/cleverdinosaur/dash](https://github.com/cleverdinosaur/dash)






<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/cleverdinosaur/dash.svg?style=for-the-badge
[contributors-url]: https://github.com/cleverdinosaur/dash/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/cleverdinosaur/dash.svg?style=for-the-badge
[forks-url]: https://github.com/cleverdinosaur/dash/network/members
[stars-shield]: https://img.shields.io/github/stars/cleverdinosaur/dash.svg?style=for-the-badge
[stars-url]: https://github.com/cleverdinosaur/dash/stargazers
[issues-shield]: https://img.shields.io/github/issues/cleverdinosaur/dash.svg?style=for-the-badge
[issues-url]: https://github.com/cleverdinosaur/dash/issues
[license-shield]: https://img.shields.io/github/license/cleverdinosaur/dash.svg?style=for-the-badge
[license-url]: https://github.com/cleverdinosaur/dash/blob/master/LICENSE
[product-screenshot]: https://cleverdinosaur.com/wp-content/uploads/2020/12/dash-768x254.png
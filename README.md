# AirCup

<div align="center">
  <img src="https://github.com/aliamerj/aircup/blob/main/client/src/assets/logo.svg" alt="AirCup">
     <h3>AirCup</h3>
  <p><strong>The Nextcloud and ownCloud Alternative: The Future of Cloud Storage</strong></p>
</div>

<br/>

## Introduction

AirCup is an open-source cloud storage platform designed to meet the increasing demands for secure, efficient, and user-friendly data management solutions. Built with speed and simplicity in mind, AirCup offers a seamless experience for businesses and individual users, featuring a one-click installation process and an intuitive user interface.

## Key Features

- **Simplified Setup:** One-click deployment, eliminating complex configuration steps.
- **Enhanced Security:** End-to-end encryption and two-factor authentication.
- **Real-Time Collaboration:** Share and work on documents in real time.
- **Superior Performance:** Efficient handling of large data transfers, reducing latency and downtime.
- **Flexible Scalability:** Dynamically adjust your storage needs and pay only for what you use.
- **Regulatory Compliance:** Built-in features to ensure compliance with global data protection regulations.

## Technology Stack

- **Frontend:** React with Vite
- **Backend:** Golang
- **Database:** SQLite

## Contributing
We welcome contributions to AirCup! If you'd like to help improve the platform, please fork the repository and create a pull request with your changes. Ensure that your code adheres to our coding standards and includes relevant tests

1. **Clone the repository**
2. **Install packages:**  ``` make install```
3. **Set up the environment variables:*
   - Copy the `.env.example` file to .env and adjust any necessary configurations.
   - Ensure your .env file is correctly set up for your development environment.
5. **Run the development environment:** ``` make dev ``` This will run both the server and the client side of the application.
6. **Compile the app and create the binary:** ``` make build ``` The compiled binary version will be located in a folder called bin in the root directory.

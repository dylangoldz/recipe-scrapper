# recipe-scrapper

## Description

This project is a simple recipe scraper that uses the SerpAPI to search for recipes and the Spoonacular API to get recipe information. It is built with Go and uses MongoDB as the database.

## Getting Started

### Prerequisites

- Go
- MongoDB
- SerpAPI
- Spoonacular API

### Installation

1. Clone the repository
2. Create a SerpAPI key:
   - Go to [SerpAPI](https://serpapi.com/) and sign up for an account
   - Once logged in, navigate to your dashboard to find your API key
3. Create a `.env` file in the root directory with the following content:
   ```
   SERP_API_KEY=your_serp_api_key_here
   MONGODB_URI=your_mongodb_connection_string_here
   ```
   Replace `your_serp_api_key_here` with your actual SerpAPI key and `your_mongodb_connection_string_here` with your MongoDB connection string.
5. Run `go run main.go` to start the server

## Email Notifications

This project includes an email notification service (`internal/notify.go`) that can send recipe updates to subscribers. Key features:

- Uses Google's SMTP server to send emails
- Requires `EMAIL_FROM` and `EMAIL_PASSWORD` environment variables
- Allows sending recipe information to specified email addresses
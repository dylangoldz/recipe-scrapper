# Recipe Web Scraper

## Description

This project is a recipe web scraper that performs daily Google searches for popular recipes, extracts standardized JSON-LD metadata from the search results, and emails the collected recipes to subscribers. It is built with Go and uses MongoDB for data storage.

## Features

- Automated daily Google searches for popular recipes using SerpAPI
- Web scraping of recipe metadata using JSON-LD format
- MongoDB integration for storing recipe data
- Email notification system for sending recipes to subscribers

## Prerequisites

- Go 1.22 or later
- MongoDB
- SerpAPI account and API key
- Gmail account for sending emails

## Installation

1. Clone the repository
2. Create a `.env` file in the root directory with the following content:
   ```
   SERP_API_KEY=your_serp_api_key_here
   MONGODB_URI=your_mongodb_connection_string_here
   EMAIL_FROM=your_gmail_address
   EMAIL_PASSWORD=your_gmail_app_password
   ```
   Replace the placeholders with your actual credentials.

3. Install dependencies:
   ```
   go mod tidy
   ```

4. Run the application:
   ```
   go run cmd/main.go
   ```

## Usage

The application will automatically perform the following tasks:

1. Search for recipes using predefined queries (see `cmd/main.go`, lines 16-37)
2. Extract recipe metadata from search results
3. Store the extracted data in MongoDB
4. Send email notifications with recipe information to subscribers

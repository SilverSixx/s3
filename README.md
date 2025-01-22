# Simple S3 Storage Manipulation Project

This project demonstrates a simple S3 storage manipulation using a combination of Vite and Tailwind CSS for the frontend, and Go for the backend.

## Project Structure

- `app/`: Contains the frontend code built with Vite and styled using Tailwind CSS.
- `api/`: Contains the backend code written in Go.

## Setup Instructions

### Frontend

1. Navigate to the `frontend` directory:
   ```sh
   cd app
   ```

2. Install dependencies:
   ```sh
   npm ci 
   ```

3. Start the development server:
   ```sh
   npm run dev
   ```

### Backend

1. Navigate to the `api` directory:
   ```sh
   cd api
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the backend server:
   ```sh
   go build -o s3 cmd/api/main.go
   chmod +x s3
   ./s3;
   ```

## Usage

1. Start both the frontend and backend servers.
2. Open your browser and navigate to the frontend server (usually `http://localhost:3000`).
3. Use the interface to interact with the S3 storage.

## License

This project is licensed under the MIT License.
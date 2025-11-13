# Examples

This directory contains examples of how to use the pingone-go-client SDK with different authentication flows.

## Available Examples

- **[authorization_code](authorization_code/)**: Interactive authentication using OAuth2 authorization code flow
- **[client_credentials](client_credentials/)**: Server-to-server authentication using client credentials grant type
- **[device_code](device_code/)**: Device authentication flow for limited-input devices or CLI applications

## Authentication Flow Guide

### Client Credentials

Best for: Server-to-server authentication, automation scripts, backend services

**PingOne Setup:**
1. Create a Worker application in your PingOne environment
2. Note the Client ID and Client Secret
3. Grant the necessary roles/permissions to the application

**Environment Variables:**
```shell
export PINGONE_CLIENT_ID=your-client-id
export PINGONE_CLIENT_SECRET=your-client-secret
export PINGONE_ENVIRONMENT_ID=your-environment-id
export PINGONE_ROOT_DOMAIN=pingone.com  # or pingone.eu, pingone.asia
```

### Authorization Code (authorization_code example)

Best for: Web applications, interactive user authentication

**PingOne Setup:**
1. Create a Native or Web application in your PingOne environment
2. Configure the **Redirect URI** to: `http://localhost:8080/callback`
   - Navigate to: Applications > [Your App] > Configuration
   - Add the redirect URI exactly as shown above (this is the SDK default)
3. Enable the authorization code grant type
4. Note the Client ID (no client secret needed for public clients)

**Environment Variables:**
```shell
export PINGONE_CLIENT_ID=your-client-id
export PINGONE_ENVIRONMENT_ID=your-environment-id
export PINGONE_ROOT_DOMAIN=pingone.com  # or pingone.eu, pingone.asia
# Optional: Customize redirect URI
export PINGONE_REDIRECT_URI_PORT=8080  # Default is 8080
export PINGONE_REDIRECT_URI_PATH=/callback  # Default is /callback
```

**What to Expect:**
- The SDK will start a local web server on port 8080
- Your browser will open to the PingOne login page
- After authentication, you'll be redirected back to the local server
- The SDK captures the authorization code and exchanges it for tokens

### Device Code (device_code example)

Best for: CLI tools, devices with limited input capabilities, headless applications

**PingOne Setup:**
1. Create a Native application in your PingOne environment
2. Enable the **device code** grant type:
   - Navigate to: Applications > [Your App] > Configuration
   - Enable "Device Authorization" grant type
3. Note the Client ID (no client secret needed)

**Environment Variables:**
```shell
export PINGONE_CLIENT_ID=your-client-id
export PINGONE_ENVIRONMENT_ID=your-environment-id
export PINGONE_ROOT_DOMAIN=pingone.com  # or pingone.eu, pingone.asia
```

**What to Expect:**
- The SDK will display a user code and verification URL
- Visit the URL in a browser on any device
- Enter the user code shown by the application
- Authenticate with your PingOne credentials
- The SDK will automatically poll and receive tokens once authentication is complete

## Running the Examples

Navigate to any example directory and run:

```shell
cd client_credentials  # client_credentials example (or authorization_code, device_code)
go run main.go
```

## Required Permissions

All examples require appropriate permissions in your PingOne environment. At minimum:
- Read access to environments (for the examples to list environments)
- Additional permissions based on your use case

Configure roles in PingOne: Applications > [Your App] > Roles

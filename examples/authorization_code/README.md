# Authorization Code Example

This example demonstrates how to use the PingOne Go Client SDK with OAuth2 Authorization Code flow for interactive user authentication.

## Use Case

Authorization Code flow is ideal for:
- Web applications with user authentication
- Applications that need user context and permissions
- Interactive authentication scenarios
- Applications requiring user consent

## PingOne Configuration

### 1. Create an Application in PingOne

1. Log in to your PingOne Admin Console
2. Navigate to **Applications**
3. Click **+ Application**
4. Select **OIDC Web App** or **Native App** depending on your use case:
   - **OIDC Web App**: For browser-based apps with a server-side component
   - **Native App**: For mobile apps or desktop applications
5. Provide a name (e.g., "Go SDK Authorization Code Example")
6. Click **Save**

### 2. Configure Redirect URI

**IMPORTANT**: The redirect URI must match what the SDK uses:

1. In your application's **Configuration** tab
2. Under **Redirect URIs**, click **+ Add**
3. Enter: `http://localhost:8080/callback`
4. Click **Save**

**Note**: This is the SDK's default. The SDK starts a local web server on port 8080 to handle the callback.

### 3. Note Your Client ID

1. On the application's **Configuration** tab, copy the **Client ID**
2. You'll use this in the environment variables below
3. **Note**: No client secret is needed for authorization code flow with public clients

## Running the Example

### Set Environment Variables

```shell
export PINGONE_CLIENT_ID="your-client-id-here"
export PINGONE_ENVIRONMENT_ID="your-environment-id-here"
export PINGONE_ROOT_DOMAIN="pingone.com"  # or pingone.eu, pingone.asia
```

Optional (to customize redirect URI):
```shell
export PINGONE_REDIRECT_URI_PORT="8080"  # Defaults to 8080
export PINGONE_REDIRECT_URI_PATH="/callback"  # Defaults to /callback
```

### Run the Example

```shell
cd examples/authorization_code
go run main.go
```

## What to Expect

1. **Local Server Starts**: The SDK starts a web server on `http://localhost:8080` (or your configured port)
2. **Browser Opens**: Your default browser opens to the PingOne login page
3. **User Authentication**: Enter your PingOne credentials
4. **Consent** (if required): Approve the requested permissions
5. **Redirect**: You're redirected back to `http://localhost:8080/callback`
6. **Token Exchange**: The SDK captures the authorization code and exchanges it for tokens
7. **API Calls**: The example makes API calls using the authenticated session
8. **Output**: Environment information is displayed in the console

## Troubleshooting

**Browser doesn't open:**
- The SDK will print the authorization URL to the console
- Manually copy and paste the URL into your browser

**"Redirect URI mismatch" error:**
- Ensure `http://localhost:8080/callback` is configured in your PingOne application
- The URI must match exactly (including the port)
- If you customized the port/path, ensure it matches your PingOne configuration

**Port 8080 already in use:**
- Stop any other services running on port 8080
- Or set `PINGONE_REDIRECT_URI_PORT` to a different port and update PingOne configuration

**"Client not found" error:**
- Verify your `PINGONE_CLIENT_ID` is correct
- Ensure the application exists in the specified environment

## Workflow Diagram

```
┌─────────────┐            ┌─────────────┐            ┌──────────────┐
│     App     │            │   Browser   │            │   PingOne    │
│    (SDK)    │            │             │            │   Server     │
└──────┬──────┘            └──────┬──────┘            └──────┬───────┘
       │                          │                          │
       │  1. Start local server   │                          │
       │     on port 8080         │                          │
       │                          │                          │
       │  2. Open browser to      │                          │
       │     authorization URL    │                          │
       ├─────────────────────────>│                          │
       │                          │                          │
       │                          │  3. Redirect to PingOne  │
       │                          │     login page           │
       │                          ├─────────────────────────>│
       │                          │                          │
       │                          │  4. Display login form   │
       │                          │<─────────────────────────┤
       │                          │                          │
       │                          │  5. User authenticates   │
       │                          ├─────────────────────────>│
       │                          │                          │
       │                          │  6. Redirect to callback │
       │                          │     with auth code       │
       │                          │<─────────────────────────┤
       │                          │                          │
       │  7. Receive callback     │                          │
       │<─────────────────────────┤                          │
       │                          │                          │
       │  8. Exchange auth code for tokens                   │
       ├────────────────────────────────────────────────────>│
       │                          │                          │
       │  9. Return: access_token, refresh_token             │
       │<────────────────────────────────────────────────────┤
       │                          │                          │
       │  10. Close local server  │                          │
       │                          │                          │
       │  11. Use tokens for API calls                       │
       ├────────────────────────────────────────────────────>│
```

## Security Notes

- This example uses `http://localhost` which is acceptable for local development
- For production applications, always use `https://` redirect URIs
- Never commit credentials or client IDs to version control
- Consider using `.env` files or secret management systems

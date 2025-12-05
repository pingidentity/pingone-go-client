# Device Code Example

This example demonstrates how to use the PingOne Go Client SDK with OAuth2 Device Code flow for CLI tools and devices with limited input capabilities.

## Use Case

Device Code flow is ideal for:
- Command-line tools and CLI applications
- Devices with limited input capabilities (smart TVs, IoT devices)
- Applications without a web browser
- Headless or remote authentication scenarios

## PingOne Configuration

### 1. Create a Native Application in PingOne

1. Log in to your PingOne Admin Console
2. Navigate to **Applications**
3. Click **+ Application**
4. Select **Native App**
5. Provide a name (e.g., "Go SDK Device Code Example")
6. Click **Save**

### 2. Enable Device Authorization Grant Type

**IMPORTANT**: You must enable the Device Authorization grant type:

1. In your application's **Configuration** tab
2. Scroll to **Grant Type**
3. Check the box for **Device Authorization** (also called Device Code)
4. Click **Save**

### 3. Note Your Client ID

1. On the application's **Configuration** tab, copy the **Client ID**
2. You'll use this in the environment variables below
3. **Note**: No client secret is needed for device code flow

## Running the Example

### Set Environment Variables

```shell
export PINGONE_CLIENT_ID="your-client-id-here"
export PINGONE_ENVIRONMENT_ID="your-environment-id-here"
export PINGONE_ROOT_DOMAIN="pingone.com"  # or pingone.eu, pingone.asia
```

### Run the Example

```shell
cd examples/device_code
go run main.go
```

## What to Expect

1. **Device Code Display**: The SDK displays:
   ```
   Please visit: https://auth.pingone.com/[environment-id]/device
   Enter code: ABCD-EFGH
   ```

2. **Visit Verification URL**: Open the URL on any device with a browser

3. **Enter Code**: Type the displayed code (e.g., `ABCD-EFGH`)

4. **User Authentication**: Log in with your PingOne credentials

5. **Approval**: Confirm the device authorization request

6. **Automatic Completion**: The SDK automatically detects successful authentication

7. **API Calls**: The example makes API calls using the authenticated session

8. **Output**: Environment information is displayed in the console

## Advantages of Device Code Flow

- **Cross-Device**: Authenticate on a phone/tablet for a CLI tool on a server
- **No Browser Required**: The CLI application doesn't need to open a browser
- **User-Friendly**: Easy-to-enter short codes instead of complex URLs
- **Secure**: No embedded secrets or credentials in the CLI tool

## Troubleshooting

**"Device authorization not enabled" error:**
- Ensure Device Authorization grant type is enabled in PingOne
- Check that you're using a Native App (not Web App or Worker)

**Code expires:**
- Device codes typically expire after 10 minutes
- Run the example again to generate a new code

**"Client not found" error:**
- Verify your `PINGONE_CLIENT_ID` is correct
- Ensure the application exists in the specified environment

**Polling timeout:**
- The SDK polls for 5 minutes by default
- If authentication isn't completed in time, the example will exit
- Run the example again to retry

## Workflow Diagram

```
┌─────────────┐                                  ┌──────────────┐
│   CLI App   │                                  │   PingOne    │
│  (SDK)      │                                  │   Server     │
└──────┬──────┘                                  └──────┬───────┘
       │                                                │
       │  1. Request device code                       │
       ├──────────────────────────────────────────────>│
       │                                                │
       │  2. Return: user_code, device_code, URL       │
       │<──────────────────────────────────────────────┤
       │                                                │
       │  3. Display code & URL to user                │
       │                                                │
       │  4. Poll for authorization                    │
       ├──────────────────────────────────────────────>│
       │                                                │
       │  5. Return: authorization_pending             │
       │<──────────────────────────────────────────────┤
       │                                                │
       
       [User visits URL on any device and enters code]
       
       │  6. Poll for authorization (repeated)         │
       ├──────────────────────────────────────────────>│
       │                                                │
       │  7. Return: access_token, refresh_token       │
       │<──────────────────────────────────────────────┤
       │                                                │
       │  8. Use tokens for API calls                  │
       ├──────────────────────────────────────────────>│
```

## Security Notes

- Device code flow is designed for public clients (no client secret)
- Codes are short-lived to prevent unauthorized access
- User must explicitly approve the device authorization
- Consider implementing additional security measures for sensitive applications

# Client Credentials Example

This example demonstrates how to use the PingOne Go Client SDK with OAuth2 Client Credentials flow for server-to-server authentication.

## Use Case

Client Credentials flow is ideal for:
- Server-to-server authentication
- Machine-to-machine (M2M) communication
- Backend services and APIs
- Automation scripts and batch jobs
- Applications that don't require user interaction

## PingOne Configuration

### 1. Create a Worker Application in PingOne

1. Log in to your PingOne Admin Console
2. Navigate to **Applications**
3. Click **+ Application**
4. Select **Worker**
5. Provide a name (e.g., "Go SDK Client Credentials Example")
6. Click **Save**

### 2. Note Your Credentials

**IMPORTANT**: Save these credentials securely - the client secret is shown only once:

1. On the application's **Configuration** tab, copy the **Client ID**
2. Copy the **Client Secret** (shown only once during creation)
3. You'll use these in the environment variables below

### 3. Assign Roles

Worker applications can have the same roles as human administrators. Grant necessary permissions:

1. Go to the **Roles** tab in your application
2. Click **Grant Roles**
3. Select the environment(s) where you want to grant access
4. Assign appropriate roles based on your needs:
   - **Environment Admin**: Full access to the environment
   - **Identity Data Admin**: User and group management
   - **Client Application Developer**: Application management
   - **Identity Data Read Only**: Read-only access to identity data
   - Or other specific roles as needed
5. Click **Save**

## Running the Example

### Set Environment Variables

```shell
export PINGONE_CLIENT_ID="your-client-id-here"
export PINGONE_CLIENT_SECRET="your-client-secret-here"
export PINGONE_ENVIRONMENT_ID="your-environment-id-here"
export PINGONE_ROOT_DOMAIN="pingone.com"  # or pingone.eu, pingone.asia
```

**Security Note**: Never commit these credentials to version control. Consider using:
- Environment variable files (`.env`) that are `.gitignore`d
- Secret management systems (AWS Secrets Manager, HashiCorp Vault, etc.)
- CI/CD platform secret storage

### Run the Example

```shell
cd examples/basic
go run main.go
```

## What to Expect

1. **Authentication**: The SDK automatically authenticates using the client credentials
2. **Token Acquisition**: An access token is obtained from PingOne
3. **API Call**: The example retrieves environment details
4. **Output**: Environment information is displayed in the console:
   ```
   Successfully read environment
   id=12345678-1234-1234-1234-123456789abc
   name=My Environment
   type=SANDBOX
   region=NA
   ```

## How It Works

The client credentials flow:
1. Application sends client ID and secret to PingOne token endpoint
2. PingOne validates the credentials
3. PingOne returns an access token
4. Application uses the token for API calls
5. Token is automatically refreshed when it expires

## Troubleshooting

**"Invalid client" error:**
- Verify your `PINGONE_CLIENT_ID` and `PINGONE_CLIENT_SECRET` are correct
- Ensure the application exists in the specified environment
- Check that the application is enabled

**"Insufficient privileges" error:**
- The application needs appropriate roles assigned
- Go to your PingOne application's Roles tab and grant necessary permissions
- Different API operations require different roles

**"Environment not found" error:**
- Verify your `PINGONE_ENVIRONMENT_ID` is correct
- Ensure the Worker App has access to this environment
- Check that roles are granted for the specific environment

**Connection errors:**
- Verify `PINGONE_ROOT_DOMAIN` matches your PingOne region
- Check your network connectivity
- Ensure firewall rules allow outbound HTTPS traffic

## Workflow Diagram

```
┌─────────────┐                                  ┌──────────────┐
│   Worker    │                                  │   PingOne    │
│     App     │                                  │   Server     │
└──────┬──────┘                                  └──────┬───────┘
       │                                                │
       │  1. Request access token                      │
       │     (client_id + client_secret)               │
       ├──────────────────────────────────────────────>│
       │                                                │
       │  2. Validate credentials                      │
       │                                                │
       │  3. Return: access_token                      │
       │<──────────────────────────────────────────────┤
       │                                                │
       │  4. Use token for API calls                   │
       ├──────────────────────────────────────────────>│
       │                                                │
       │  5. Return: API response                      │
       │<──────────────────────────────────────────────┤
       │                                                │
       │  6. Token expires, request new token          │
       │     (automatic refresh)                       │
       ├──────────────────────────────────────────────>│
       │                                                │
       │  7. Return: new access_token                  │
       │<──────────────────────────────────────────────┤
```

## Security Best Practices

1. **Protect Client Secret**: Treat it like a password - never commit to source control
2. **Rotate Credentials**: Regularly rotate client secrets
3. **Least Privilege**: Grant only the minimum required roles
4. **Secure Storage**: Use secure secret management systems in production
5. **Audit Access**: Monitor API usage through PingOne audit logs
6. **Environment Isolation**: Use separate Worker Apps for dev/test/prod environments

## Additional Resources

- [PingOne Worker Apps Documentation](https://docs.pingidentity.com/r/en-us/pingone/p1_add_app_worker)
- [OAuth2 Client Credentials Flow](https://oauth.net/2/grant-types/client-credentials/)
- [PingOne API Documentation](https://apidocs.pingidentity.com/pingone/platform/v1/api/)

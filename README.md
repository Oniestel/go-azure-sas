# go-azure-sas
Go package for generating Azure Shared Access Signatures (SAS) token.

## Example
```
token := azuresas.GenerateToken("https://namespace.servicebus.windows.net/default", "keyName", "key", 3600)
```

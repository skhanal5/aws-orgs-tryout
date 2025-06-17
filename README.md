# aws-orgs-tryout
Experimenting with AWS Organizations, specifically cross-account roles and parent-child relationships with IAM.

## Configuration

Suppose you have the following OU Path in AWS Organizations:

```
Root
├── Prod
│   ├── Prod Account
│   └── Target Account
└── Management Account
```

where `Prod Account` contains an IAM Role granting any principal to assume into a role in `Target Account`. An example of both can be found in the `iam` directory. 

## Cross-Account Role Example

### Flow Diagram

```
Resource (e.g., EC2 instance)
   │
   └─ Has an IAM execution role: {{PROD_EXECUTION_ROLE}}
         │
         ▼
  {{PROD_EXECUTION_ROLE}} is trusted by PROD_ACCOUNT and
  allowed to assume AssumeMemberRole in PROD_ACCOUNT
         │
         ▼
  Program running on resource uses {{PROD_EXECUTION_ROLE}} creds
  → Calls STS:AssumeRole on AssumeMemberRole (in PROD_ACCOUNT)
         │
         ▼
  Gets temporary creds for AssumeMemberRole
         │
         ▼
  Uses those creds to call STS:AssumeRole on PrivilegedMemberRole (in MEMBER_ACCOUNT)
         │
         ▼
  Gets temporary creds for PrivilegedMemberRole
         │
         ▼
  Uses PrivilegedMemberRole creds to perform operations in MEMBER_ACCOUNT
```
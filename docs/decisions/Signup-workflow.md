# Signup Workflows

There are two possible workflows. An individual or an organization. The intention of the workflow is onboard a person or a team in as few clicks as necessary while addressing a few business issues:

1. Security First
1. Cut down on spam and abuse
1. We get paid

## Individual

### Security First

Early prototype won't enforce 2FA, but before a bigger public release, this should exist. I don't want to tie anyone to a specific OAuth provider, we'll use the appropriate library to generate the QR code and TOTP for the likes of 1Password and what not.

### Cut down on spam and abuse

The easiest thing to do is e-mail verification. This will exist from the beginning. I realize it may impact conversion or signup, but it's too easy to use some BS email account and abuse a system. I'll take fewer good people than a lot of bad people. Another possibility is Snipe's service for email verification.

### We get Paid

This will help with spam long-term, but we're not running a traditional startup where the first one is free and folks can upgrade. We'll do a cost analysis of the system and ensure margins are big enough and we're not running on a shoe-string budget and we can pay folks a living wage. While I would love to give back to the community and give open source libraries free access, they can download the code and run it themselves.

## Organizational

### Security First

Like individuals, the person signing up is the org admin to start with. They can later delegate or share that role.

### Cut down on spam and abuse

Like individuals, we'll also do e-mail verifications. In addition to org admin, we'll also provide a mechanism for billing admin as well as other individuals in the org. We'll verify these contact addresses as well. We'll probably do some e-mail verification for org name vs domain name and/or DNS TXT record. And should an org want to change billing admin, we'll probably require some e-mail verification from org admin.

### We get Paid

They have money, we we're providing a service that we want paid for. Business 101. Billing will be tied to account operation. Keep the machine fed with money, it continues to operate. Billing fails, notifications are sent and machine stops working. We'll always give folks a means to download their data, but we're not operating a charity here.

## Backend Things

### Individual

When an individual signs up, we ask for email address (which becomes their username), this is unique, and we'll parse `+ aliases` so that folks can't work around spam prevention/account limits or what not. And we'll ask for a password. Perhaps a library that shows strength/quality, but don't limit to some arbitrary 8, 16, 20 character bullshit. We'll see `verified` to `false`, send an email notification forcing e-mail verification and upon email verification, `verified` will be set to `true`.

We will do an N number day/week trial to allow folks to take things for a spin. We will _not_ collect payment information until the trial expires. They will have access to their data after the trial expires as it is their data.

### Organziation

When a user creates an org, a few things need to happen:

1. Create the org super user (this is the user creating the org initially).
1. Create the org.
1. Associate the user with the org.

Much like the individual workflow, we will want to verify the email address. We can take a step further and parse the email to ensure the tld matches the org name or require a DNS TXT record that shows they are in control of that org.

Much like an individual, we will do an N number day/week trial to allow them and their team to take the product for a spin. We will _not_ collect payment information

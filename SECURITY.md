# Security Policy

## Reporting a Vulnerability

The Base-FullStack-admin team takes security vulnerabilities seriously. We appreciate your efforts to responsibly disclose your findings and will make every effort to acknowledge your contributions.

To report a security vulnerability, please email us at [mo7amed4522@gmail.com](mailto:mo7amed4522@gmail.com) with the subject line "[SECURITY] Tourism-FullStack-admin Vulnerability Report". You can also contact us via WhatsApp at [+971 052 537 4187](https://wa.me/971525374187) for urgent matters.

Please include the following information in your report:

- Description of the vulnerability
- Steps to reproduce the issue
- Potential impact of the vulnerability
- Any suggestions for addressing the vulnerability
- Your name/handle (if you wish to be credited)

We will acknowledge receipt of your vulnerability report within 48 hours and will send you regular updates about our progress.

## Security Measures

Tourism-FullStack-admin implements the following security measures to protect user data and system integrity:

### Authentication & Authorization

- **JWT-based Authentication**: Secure token-based authentication with proper expiration and refresh mechanisms
- **Role Separation**: Different JWT tokens for different user roles (admin, agent, customer, guest)
- **Permission-based Access Control**: Fine-grained RBAC system using Casbin
- **Mobile Authentication**: Secure OTP verification for mobile users
- **Token Lifecycle Management**: Proper token rotation, blacklisting, and revocation

### Data Protection

- **Encryption**: Sensitive data is encrypted at rest and in transit
- **Payment Security**: No storage of CVV/CVC codes, compliance with PCI DSS guidelines
- **Input Validation**: Comprehensive validation of all user inputs
- **Output Encoding**: Proper encoding of data to prevent XSS attacks
- **CSRF Protection**: Implementation of anti-CSRF tokens

### Infrastructure Security

- **TLS/SSL**: All communications are encrypted using TLS 1.3
- **Rate Limiting**: Protection against brute force and DDoS attacks
- **Dependency Scanning**: Regular scanning and updating of dependencies
- **Secure File Storage**: Proper access controls for uploaded files
- **Database Security**: Parameterized queries to prevent SQL injection

### Monitoring & Auditing

- **Logging**: Comprehensive logging of security-relevant events
- **Audit Trails**: Tracking of all administrative actions
- **Intrusion Detection**: Monitoring for suspicious activities
- **Regular Security Reviews**: Scheduled security assessments

## Vulnerability Disclosure Timeline

We follow these general timelines for addressing security vulnerabilities:

1. **Acknowledgment**: Within 48 hours of receiving a report
2. **Validation**: Within 1 week of receiving a report
3. **Remediation Plan**: Within 2 weeks of validation
4. **Fix Implementation**: Timeline depends on severity:
   - Critical: Within 7 days
   - High: Within 14 days
   - Medium: Within 30 days
   - Low: Within 90 days
5. **Public Disclosure**: Coordinated with the reporter after the fix has been deployed

## Supported Versions

Only the latest major version of Tourism-FullStack-admin receives security updates. We encourage all users to keep their installations up to date.

## Security Best Practices for Deployment

When deploying Tourism-FullStack-admin, we recommend the following security best practices:

1. **Environment Configuration**: Never store sensitive information in code or version control
2. **Database Security**: Use strong passwords and restrict database access
3. **Regular Updates**: Keep all dependencies and the application itself updated
4. **Firewall Configuration**: Properly configure firewalls to restrict unnecessary access
5. **Monitoring**: Implement monitoring for unusual activities
6. **Backups**: Regularly backup data and test restoration procedures
7. **Access Control**: Implement the principle of least privilege for all accounts

## Contact

For any security-related questions or concerns, please contact:

- **Name:** Khaled Mohamed
- **Email:** [mo7amed4522@gmail.com](mailto:mo7amed4522@gmail.com)
- **Phone:** +971 052 537 4187 (UAE)
- **WhatsApp:** [+971 052 537 4187](https://wa.me/971525374187)

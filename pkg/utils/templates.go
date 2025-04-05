package utils

const otpEmailTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Your Verification Code</title>
    <style>
        body, p, h1, h2 {
            font-family: 'Poppins', Arial, sans-serif;
        }
    </style>
</head>
<body style="margin:0;padding:0;background:#f4f7fa;font-family:'Poppins',sans-serif;">
    <table width="100%" style="max-width:600px;margin:auto;background:#fff;">
        <tr>
            <td style="background:#4169e1;padding:25px;text-align:center;">
                <h1 style="color:#fff;margin:0;font-weight:600;font-size:24px;">VithsutraTechnologies</h1>
            </td>
        </tr>
        <tr>
            <td style="padding:40px 30px 20px;text-align:center;">
                <h2 style="color:#333;font-weight:600;font-size:22px;">Your Verification Code</h2>
                <p style="color:#555;text-align:left;">
                    We received a request to verify your identity. Use the code below to complete the process:
                </p>
                <div style="background:#f8f9fa;border:1px solid #e9ecef;border-radius:6px;padding:20px;margin:30px auto;max-width:300px;">
                    <h1 style="font-family:'Courier New',monospace;letter-spacing:5px;color:#4169e1;font-size:32px;">{{.OTP}}</h1>
                </div>
                <p style="color:#555;text-align:left;">
                    <strong>Important:</strong> This code will expire in <span style="color:#4169e1;font-weight:600;">{{.ExpireMinutes}} minutes</span>.
                </p>
                <p style="color:#555;text-align:left;">
                    If you did not request this code, please ignore this email or contact our support team.
                </p>
            </td>
        </tr>
        <tr>
            <td style="padding:0 30px 30px;">
                <div style="background:#f8f9fa;border-left:4px solid #4169e1;padding:15px;border-radius:0 4px 4px 0;">
                    <h3 style="margin:0 0 10px;color:#333;font-size:16px;font-weight:600;">Security Reminder</h3>
                    <p style="margin:0;color:#555;font-size:14px;">
                        • Never share this code with anyone.<br>
                        • VithsutraTechnologies will never ask for your password.<br>
                        • Ensure you're on our official website before entering credentials.
                    </p>
                </div>
            </td>
        </tr>
        <tr>
            <td align="center" style="background:#f8f9fa;padding:20px 30px;color:#666;font-size:14px;border-top:1px solid #eee;">
                <p style="margin:0 0 10px;">© 2025 VithsutraTechnologies. All rights reserved.</p>
                <p style="margin:0;">
                    Need help? Contact us at <a href="mailto:support@vithsutratechnologies.com" style="color:#4169e1;">support@vithsutratechnologies.com</a>
                </p>
            </td>
        </tr>
    </table>
</body>
</html>
`

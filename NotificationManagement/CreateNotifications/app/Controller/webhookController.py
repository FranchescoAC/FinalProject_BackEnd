import smtplib
from email.mime.text import MIMEText
from flask import jsonify, request
from ..Model.createModel import create_notification

# Receive webhook and create a new notification, then send an email
def receive_webhook():
    try:
        # Get the data sent by UserManagement
        data = request.get_json()

        # Validate that the required fields are present
        if not data or "user_id" not in data or "username" not in data or "email" not in data:
            return jsonify({"error": "Invalid payload"}), 400

        # Prepare the notification data to store
        notification_data = {
            "user_id": data["user_id"],
            "username": data["username"],
            "email": data["email"],
            "message": f"Welcome {data['username']}! Your account has been successfully created.",
            "status": "sent"
        }

        # Call the model function to store the notification
        notification_id = create_notification(notification_data)

        # Send the welcome email
        send_email(
            recipient=data["email"],
            subject="Welcome to Our Service of Travel in Ecuador!",
            body=f"Hi {data['username']},\n\nThis is a confirmation email that your account has been successfully created, thank you for choosing us.\n\nBest regards,\nYour Company"
        )

        return jsonify({"message": "Notification created and email sent", "id": str(notification_id)}), 201

    except Exception as e:
        return jsonify({"error": f"An error occurred: {str(e)}"}), 500

# Function to send an email using SMTP
def send_email(recipient, subject, body):
    sender_email = "vinculacionvcs@gmail.com"  # Replace with your sender email
    sender_password = "aklw zuug wgfv ebso"  # Replace with your sender email password

    # Setup the MIMEText object with the email content
    message = MIMEText(body)
    message["Subject"] = subject
    message["From"] = sender_email
    message["To"] = recipient

    try:
        # Connect to the SMTP server (example uses Gmail)
        with smtplib.SMTP("smtp.gmail.com", 587) as server:
            server.starttls()  # Secure the connection
            server.login(sender_email, sender_password)
            server.sendmail(sender_email, recipient, message.as_string())
            print(f"✅ Email sent to {recipient}")

    except Exception as e:
        print(f"❌ Failed to send email: {str(e)}")

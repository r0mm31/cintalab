import os
import json
import requests

def send_slack_notification(message):
    webhook_url = os.getenv('SLACK_WEBHOOK_URL')
    payload = {
        'text': message
    }
    
    response = requests.post(webhook_url, data=json.dumps(payload), headers={'Content-Type': 'application/json'})
    
    if response.status_code != 200:
        raise ValueError(f'Request to Slack returned an error {response.status_code}, the response is:\n{response.text}')

if __name__ == "__main__":
    message = "La imagen de Alex y Roman se ha subido correctamente a Docker Hub."
    send_slack_notification(message)

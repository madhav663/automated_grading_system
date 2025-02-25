import requests

OLLAMA_API_URL = "http://ollama-api-url"  
OLLAMA_API_KEY = "your-ollama-api-key"   

def get_feedback(text):
    headers = {
        "Authorization": f"Bearer {OLLAMA_API_KEY}",
        "Content-Type": "application/json"
    }
    
    data = {
        "text": text,
        "task": "feedback_generation"  
    }

    response = requests.post(f"{OLLAMA_API_URL}/generate", json=data, headers=headers)
    
    if response.status_code == 200:
        result = response.json()
        return result.get("feedback", "No feedback generated")
    else:
        return "Error generating feedback"


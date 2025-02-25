import requests

OLLAMA_API_URL = "http://ollama-api-url"  
OLLAMA_API_KEY = "your-ollama-api-key"   

def check_phishing(text):
    headers = {
        "Authorization": f"Bearer {OLLAMA_API_KEY}",
        "Content-Type": "application/json"
    }
    
    data = {
        "text": text,
        "task": "phishing_check"  
    }

    response = requests.post(f"{OLLAMA_API_URL}/check", json=data, headers=headers)
    
    if response.status_code == 200:
        result = response.json()
        return {"phishing": result.get("phishing", False)}
    else:
        return {"phishing": False, "error": "Error communicating with Ollama API"}

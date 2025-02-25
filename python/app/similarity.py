import requests

OLLAMA_API_URL = "http://ollama-api-url"  
OLLAMA_API_KEY = "your-ollama-api-key"   

def check_similarity(text1, text2):
    headers = {
        "Authorization": f"Bearer {OLLAMA_API_KEY}",
        "Content-Type": "application/json"
    }
    
    data = {
        "text1": text1,
        "text2": text2,
        "task": "similarity_check" 
    }

    response = requests.post(f"{OLLAMA_API_URL}/check", json=data, headers=headers)
    
    if response.status_code == 200:
        result = response.json()
        return result.get("similarity_score", 0)
    else:
        return 0  

from flask import Flask, request, jsonify
from phishing import check_phishing
from similarity import check_similarity
from feedback import get_feedback

app = Flask(__name__)

@app.route('/check_phishing', methods=['POST'])
def phishing():
    data = request.json
    text = data['text']
    
   
    result = check_phishing(text)
    return jsonify(result)

@app.route('/check_similarity', methods=['POST'])
def similarity():
    data = request.json
    text1 = data['text1']
    text2 = data['text2']
    
   
    similarity_score = check_similarity(text1, text2)
    return jsonify({"similarity_score": similarity_score})

@app.route('/generate_feedback', methods=['POST'])
def feedback():
    data = request.json
    text = data['text']
    
   
    feedback_result = get_feedback(text)
    return jsonify({"feedback": feedback_result})

if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0", port=5000)

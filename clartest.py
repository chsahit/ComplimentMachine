from clarifai.rest import ClarifaiApp

app = ClarifaiApp("ixACIQvGqKKcJCGLT_xnEh4_jlG7dRKXuzF4jam3", "-XQ5gLtB0ZljTUEmoaqV4LI8UXdlwZNEvLTkSXt-")
      
model = app.models.get('general-v1.3')

response = model.predict_by_url(url = 'https://avatars1.githubusercontent.com/u/3252741?v=3&s=400')

for tag in response["outputs"][0]["data"]["concepts"]:
    print(tag["name"])

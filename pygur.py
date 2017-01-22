import pyimgur
import sys

im = pyimgur.Imgur('37a3379f24710f8')
uploaded_image = im.upload_image(sys.argv[1], title="")
#print(uploaded_image.title)
print(uploaded_image.link)
#print(uploaded_image.size)
#print(uploaded_image.type)
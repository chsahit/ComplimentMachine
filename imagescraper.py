import image_scraper
import subprocess
searchterm = "cat"
URL = "https://www.google.com/search?q="+searchterm+"&espv=2&source=lnms&tbm=isch&sa=X&ved=0ahUKEwjei7XD2tPRAhVI4CYKHSmxBgwQ_AUICCgB&biw=1440&bih=826"
subprocess.popen(["python image_scraper"," -s pics". "-m 2", URL])

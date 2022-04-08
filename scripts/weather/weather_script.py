import requests
from requests.structures import CaseInsensitiveDict
import geocoder
import urllib.request
def isOnline(host='http://google.com'):
    try:
        urllib.request.urlopen(host) #Python 3.x
        return True
    except:
        return False

iconDict = {

	"01":"☀️",
	"02":"⛅",
	"03":"🌥️",
	"04":"☁️",
	"09":"🌧️",
	"10":"🌦️",
	"11":"⛈️",
	"13":"❄️",
	"50":"🌫️"
 }

def main():

	if isOnline() == False:
		print("computer offline")
		return

	try:

		g = geocoder.ip('me').latlng
		lat = str(g[0])
		lon = str(g[1])

		#print( g )

		url = "http://api.openweathermap.org/data/2.5/weather?lat="+lat+"&lon="+lon+"&appid=4c1ae3c1a2cbd1c3ff74c66b9305557a&units=metric"

		headers = CaseInsensitiveDict()
		headers["Authorization"] = "Bearer 4c1ae3c1a2cbd1c3ff74c66b9305557a"

		resp = requests.get(url, headers=headers)

		respCode = resp.status_code

		if (respCode == 200):
			respData = resp.json()

			#print(respData)

			iconKey = str(respData['weather'][0]['icon'])[:2]

			if ( iconKey in iconDict ):
				print( iconDict[iconKey], end=" " )
			else:
				print( f"[{iconKey}]", end=" "  )

			print( str(respData['main']['temp'])[:4] + "°C")
			return
		print(f"Error: {respCode}")
	except Exception as e:
		print( "error" )

if __name__ == "__main__":
	main()

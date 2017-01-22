<!DOCTYPE html>
<html>
<title>Compliment Machine</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="http://www.w3schools.com/lib/w3.css">
<body class="w3-content" style="max-width:1500px">

<!-- First Grid: Logo & About -->
<div class="w3-row">
  <div class="w3-half w3-red w3-hide-small w3-container w3-center" style="height:94px " >
  <div class="w3-half w3-red w3-hide-large w3-hide-medium w3-container w3-center" style="height:114px " >

    <div class=" w3-center style" style = "height:300px ">
      <h1 class = "w3-center">What are your best qualities?</h1>
    </div>
    <div class = "w3-row w3-margin ">
        <img src="{{.image}}">
    </div>
    <div class = "w3-row w3-margin " style="height: 25% width: 25%">
        <img src="{{.image}}">
    </div>
    <div class="w3-padding-64">
    </div>
  </div>
  <div class=" w3-red w3-container" style="height:94px">
    <div class="w3-padding-64 w3-center">

      <div class=" w3-padding-xxlarge" >
        <div class='w3-panel w3-blue w3-round-xlarge w3-padding-16' >
            <p style='color: white;'> {{ (index .tags 0).Trait }} </p>
        </div>
        <div class='w3-panel w3-blue w3-round-xlarge w3-padding-16' >
            <p style='color: white;'> {{ (index .tags 1).Trait }} </p>
        </div>
        <div class='w3-panel w3-blue w3-round-xlarge w3-padding-16' >
            <p style='color: white;'> {{ (index .tags 2).Trait }} </p>
        </div>
      </div>
    <h style="color:black;"> These are some trending words of people around you! </h>
    <!--<div class='w3-panel w3-blue w3-round-xlarge w3-padding-16' >
        <p style='color: white;'> {{ (index .shared_tags 0).Trait }} </p>
    </div>
    <div class='w3-panel w3-blue w3-round-xlarge w3-padding-16' >
        <p style='color: white;'> {{ (index .shared_tags 1).Trait }} </p>
    </div>
    <div class='w3-panel w3-blue w3-round-xlarge w3-padding-16' >
        <p style='color: white;'> {{ (index .shared_tags 2).Trait }} </p>
    </div>-->
  </div>

    </div>
  </div>
</div>

</body>
</html>

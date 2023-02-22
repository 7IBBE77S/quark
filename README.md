# Project website
### Run command: 'go run cmd/web/*.go (NOTE: NOT go run *.go)or go run cmd/web/main.go middleware.go routes.go'
- Built with Go version 1.20
- Dependencies:
    * Chi router
    * Nosurf middleware
    * SCS session
- File note: 
    * .page.tmpl functions as an .html file similar to .blade.php, .ejs, .cshtml, etc.
    * .layout.tmpl functions as a layout file similar to layout/master.blade.php (e.g. @extends("layout.master"), shared/filename.ejs, (e.g. <%-include("../shared/header");%>), Shared/_Layout.cshtml (e.g. @{ViewData["Title"] = "Index";}), etc.

## Website project file structure explanation
- cmd/web
	* main.go
	* routes.go
	* middleware.go
	* handlers.go

## TODO:

### EX01_02
- [?] Describe the product
- [?] Mock screenshots
- [ ] Team bios
- [✓] Two web pages

### EX01_02
- [?] Document the project
- [?] Describe the problem
- [?] Explain the target audience
- [?] Device requirements
- [ ] System requirements
  
### EX01_03
- [ ] Pitch the idea
- [ ] Interview five people
- [ ] Receive feedback

### Bugs
- [ ] Fix the hamburger button not showing up correctly on mobile

### Stretch Goals
- [ ] Remove dependencies (create a basic router, session, and middleware)
- [ ] Add a button that generates a 2D random noise map in a window or as an image.

### Pseudo-code for potential noise generation 🤷‍♂️
```go

type Noise interface 
	two x, y  float 64 bit
	three x, y, z float 64 bit
	four x, y, z, w  float 64 bit

func seed int 64 bit  Noise 
	s = &noise
	permutation	  make int 16 bit array, 256,
	permutation3D make int 16 bit array, 256,
	

	source = make int 16 bit array, 256,
	loop through the range of source
	source[i] = int16(i)
	

	seed = seed*123 + 123
	seed = seed*123 + 123
	seed = seed*123 + 123

loop through the range of 255 to 0 decrement i
	seed = seed*123 + 123
    j = int 32 bit, seed + 31 % int 64 bit i + 1

loop through the range of 255 to 0 decrement i
	
	seed = seed*123 + 123
    j = int 32 bit, seed + 31 % int 64 bit i + 1
    if j is less than 0 then increment j by i + 1

	s.permutation i = source j
    s.permutation3D i = s.permutation i mod int 16 bit length(noise gradient) divided by 3 multiplied by 3
    source j = source i
	

	return s

```

OR, from a series https://www.youtube.com/watch?v=WP-Bm65Q-1Y&list=PLFt_AvWsXl0eBW2EiBtl_sxmDtSgZBxB3&index=2
...the source from https://github.com/SebLague/Procedural-Landmass-Generation/blob/master/Proc%20Gen%20E02/Assets/Scripts/Noise.cs

```c#
using UnityEngine;
using System.Collections;

public static class Noise {

	public static float[,] GenerateNoiseMap(int mapWidth, int mapHeight, float scale) {
		float[,] noiseMap = new float[mapWidth,mapHeight];

		if (scale <= 0) {
			scale = 0.0001f;
		}

		for (int y = 0; y < mapHeight; y++) {
			for (int x = 0; x < mapWidth; x++) {
				float sampleX = x / scale;
				float sampleY = y / scale;

				float perlinValue = Mathf.PerlinNoise (sampleX, sampleY);
				noiseMap [x, y] = perlinValue;
			}
		}

		return noiseMap;
	}

}
```
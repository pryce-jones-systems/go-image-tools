# go-image-tools
This library provides image manipulation tools for grayscale images, intended for computer vision and machine learning applications. I could add support for colour images, but havn't found the need yet, as most computer vision pipelines operate on grayscale images. Drop me an email if you want me to add colour image support (enquiries@pryce-jones-systems.com).

The library is designed to take advantage of Go's fantastic concurrency features. As a result the code may at times look a bit convoluted and difficult to follow, but it does provide a great speed boost on machines with lots of CPU cores. I've done my best to comment everything, so it shouldn't be too bad.

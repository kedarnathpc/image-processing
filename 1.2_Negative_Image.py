from PIL import Image
import numpy as np
import matplotlib.pyplot as plt

image_path = "image.jpeg"
color_image = Image.open(image_path)
color_image_array = np.array(color_image)
height, width, channels = color_image_array.shape

negative_image = np.empty_like(color_image_array)

grayscale_image = np.empty_like(color_image_array)
grayscale_negative_image = np.empty_like(color_image_array)

for i in range(height):
    for j in range(width):
        for k in range(channels):
            negative_image[i, j, k] = np.clip(255 - color_image_array[i, j, k], 0, 255)

for i in range(height):
    for j in range(width):
        r, g, b = color_image_array[i, j]
        grayscale_image[i, j] = 0.2989 * r + 0.5870 * g + 0.1140 * b
        grayscale_negative_image[i, j] = 255 - grayscale_image[i, j]

fig, axs = plt.subplots(2, 2)
axs = axs.flatten()

image_arrays = [color_image_array, negative_image, grayscale_image, grayscale_negative_image]
titles = ["Color Image", "Color Negative Image", "Grayscale Image", "Grayscale Negative Image"]

for i, ax in enumerate(axs):
    ax.imshow(image_arrays[i].astype(np.uint8))
    ax.axis('off')
    ax.set_title(titles[i])

plt.tight_layout()
plt.show()


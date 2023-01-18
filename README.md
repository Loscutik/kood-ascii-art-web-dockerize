# Project name: Ascii-art-dockerize

The project ASCII-ART-WEB-DOCKERIZE uses Docker to run a server to host a website where you can write a custom text in ASCII art style.
<br>
<br>
<br>

## Usage:

### Part 1. Launch docker and open the website. 

0. [Install Docker](https://docs.docker.com/engine/install/) if you havn't got it yet. [There](https://stackoverflow.com/questions/48957195/how-to-fix-docker-got-permission-denied-issue) is an hint how to make Docker run without sudo on Linux. 

1. Run the script dockerize.sh to build the docker image and start the container. 

2. Visit the website by the link: http://localhost:8080/

 > [!Note]
 > If you want to run a command `docker exec -it <container_name> /bin/bash` from audit description use `/bin/sh` instead of `/bin/bash`, 
 > because a small alpine base image was used for this image which does not contain `bash`

<br>
<br>
<br>

### Part 2. Use the website.

On the opened website, you can see the field "Type your text here". Please type there the text, that you want to write with ASCII symbols.

Then, you can choose one of the font styles: standard, shadow, or thinkertoy. These are different ASCII fonts for your text.

After that, you can choose a color of your text and background.

Then, please click "Submit", scroll down, and enjoy the results 😉

<br>

*NB! If your final ASCII result doesn't fit in the field, please just put your mouse to the top bottom corner of the result field, and stretch it.*

<br>



We hope you will like it as we do 🥰!

<br>
<br>
<br>

## Authors

Created by: Olena Budarahina (Gitea username: obudarah), Kristina Volkova (Gitea username: Mustkass).


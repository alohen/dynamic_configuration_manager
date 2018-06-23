# dynamic_configuration_manager

this prgoram is a go server which dynamiclly builds GUI to edit configuration,
according to it's Go struct fields. 

!!this is only an alpha, and most things are not suppoerted yet!!

For Example, for the following struct:

![alt text](https://image.ibb.co/gZxZqT/struct.png)

And The Following config: 

![alt text](https://image.ibb.co/k1WG4o/config.png)

If we request this config, which is located under the config floder, in the path: "persons/alon" we will receive: 

![alt text](https://image.ibb.co/k34sc8/ui.png)

As seen in the GUI image, the numerical fields are already validated on client side, based only on the 
face that the type of the field in the struct is numeric.

Features soon to come: 
* editability of a struct field controlled by a tag 
* dynamic building of client side validation according to struct tags
* support of lists 
* support of inner structs

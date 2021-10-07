
CREATE TABLE OLIMPIC(
	idOlimpic INT identity(1,1) NOT NULL,
	nombre CHAR(50),
	comentario CHAR(300),
	fecha DATE,
	hashtags VARCHAR(300),
	upvotes INT,
	downvotes INT,
	PRIMARY KEY(idOlimpic)
);

SELECT * FROM OLIMPIC;

DROP TABLE OLIMPIC

SET DATEFORMAT dmy;  
GO

INSERT INTO OLIMPIC(nombre, comentario, fecha, hashtags, upvotes, downvotes)
VALUES('Leonel Aguilar', 'el dia de hoy el atleta juan perez logro oro', '24/07/2021', 'remo,atletismo,natacion', 100, 30);
INSERT INTO OLIMPIC(nombre, comentario, fecha, hashtags, upvotes, downvotes)
VALUES('Juan Carlos', 'Tiro y gano', '25/07/2021', 'tiro,arco,tiroconarco', 100, 30);

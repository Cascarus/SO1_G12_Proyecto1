use serde::{Deserialize, Serialize};
use chrono::{ NaiveDate};
//use futures::executor::block_on;
   
/*
    se coloca Serializacion y Deserializacion para poder convertir en Formato Json
    y viceversa.
*/

#[derive(Debug,Deserialize,Serialize)]
pub struct Tuits{
  pub  nombre: String,
  pub  comentario: String,
  pub  fecha: String,
  pub  hashtags: Vec<String>,
  pub  upvotes: i32,
  pub  downvotes: i32
}

#[derive(Debug,Deserialize,Serialize,PartialEq, Eq)]
pub struct Notificacion{
   pub guardados: i64,
   pub api: String,
   pub tiempo: u64,
   pub db: String,
}

#[derive(Debug,Deserialize,Serialize,PartialEq, Eq)]
pub struct Tuit{
   pub nombre: Option<String>,
   pub comentario: Option<String>,
   pub fecha: NaiveDate,
   pub hashtags: Option<String>,
   pub upvotes: i32,
   pub downvotes: i32
}


#[derive(Debug,Deserialize,Serialize)]
pub struct Mensaje{
  pub  mensaje: String,
}


pub fn concatenarcadenas(lista:Vec<String>)->String{
    let mut indice = 0;
    let mut cadena_concatenada = String::new();
    while indice < lista.len() {
           let cadena:String = lista[indice].to_string();
           if indice == lista.len()-1 {
             cadena_concatenada += &cadena;      
           }else{
              let coma:String = ",".to_string();    
              cadena_concatenada += &cadena;      
              cadena_concatenada += &coma;      
           }
           indice += 1;
    }
    return format!("{}",cadena_concatenada); 
}



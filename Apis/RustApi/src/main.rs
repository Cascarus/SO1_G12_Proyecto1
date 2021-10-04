mod publishmessage;
// Metodos
use crate::publishmessage::publicar_mensaje ;
use crate::publishmessage::concatenarcadenas ;
// Struct
use crate::publishmessage::Mensaje;
use crate::publishmessage::Tuits;
use crate::publishmessage::Tuit;
use crate::publishmessage::Notificacion;

use axum::{
    handler::{get, post},
    response::IntoResponse,
    Json, Router
};
use mysql::prelude::*;
use mysql::*;
//use chrono::{DateTime, NaiveDate, NaiveDateTime, NaiveTime};
use chrono::{ NaiveDate};

//use chrono::format::ParseError;
use mongodb::{Client, options::ClientOptions};
use std::time::{Instant, Duration};
use std::thread;
//use mongodb::bson::{doc, Document};

// use mongodb::{error::Error, Collection};
// use mongodb::results::{  InsertOneResult};

use std::net::SocketAddr;
// use std::mem;
use dotenv::dotenv;
use std::env;
//use crate::publishmessage::Notificacion;
static mut CONTADORCOSMODB : i64 = 0;
static mut CONTADORSQLDB: i64 = 0;
static mut CARGAR: bool = false;
static mut SEGUNDOSMYSQL: u64 = 0;
static mut SEGUNDOSMONGO: u64 = 0;


#[tokio::main]
pub async fn main() {

    dotenv().ok();
    let app = Router::new()
           .route("/iniciarcarga/rust/", get(iniciar_cargar))
           .route("/publicar/rust/", post(post_publicar_carga))
           .route("/finalizarcarga/rust/", post(finalizar_carga));

    // run our app with hyper
    // `axum::Server` is a re-export of `hyper::Server`
    let addr = SocketAddr::from(([0, 0, 0, 0], 4000));
    println!("listening on {}", addr);
    axum::Server::bind(&addr)
    .serve(app.into_make_service())
    .await
    .unwrap();
   
}

pub async fn iniciar_cargar()->Json<Mensaje>{
    unsafe {
      if CARGAR == false{
             CARGAR = true;    
             CONTADORCOSMODB = 0;
             CONTADORSQLDB = 0;
             let smsjson = Mensaje { mensaje: "Se ha realizado la conexion exitosamente".to_string() };
             return Json(smsjson); 
      };
    }
    let smsjson = Mensaje { mensaje: "Actualmente estas Conectado!".to_string() };
    Json(smsjson)  
}

pub async fn post_publicar_carga(Json(_req): Json<Tuits>)-> impl IntoResponse {
         
    let database_url = env::var("DATABASE_URL").expect("DATABASE URL is not in .env file");
    let client_options = ClientOptions::parse(&database_url).await.unwrap();
    let client = Client::with_options(client_options).unwrap();
    let db = client.database("Olympics");
    let collection = db.collection::<Tuits>("Tuits");      
    unsafe{
        // let mut arreglo = vec![];
                       let  _tuiteo = Tuits {
                            nombre: _req.nombre.to_string(),
                            comentario: _req.comentario.to_string(),
                            fecha: _req.fecha.to_string(),
                            hashtags:_req.hashtags.to_vec(),
                            upvotes: _req.upvotes,
                            downvotes: _req.downvotes
                          };
          //arreglo.push(_tuiteo);
          collection.insert_one(_tuiteo, None).await.unwrap();
            
          
          println!("Se inicio La carga en CosmoDB");
          CONTADORCOSMODB+=1;
          let start = Instant::now();
          expensive_function();
          let duration = start.elapsed();
            SEGUNDOSMONGO = duration.as_secs();

         
    //println!("Time elapsed in expensive_function() is: {:?}", duration); 

    }          
    
    

    let url = "mysql://root:123456@34.122.151.115/Olympics";
    let opts = Opts::from_url(url).unwrap();
    let pool = Pool::new(opts).unwrap();
    let mut conn = pool.get_conn().unwrap(); 
    unsafe{

      let mut insert_data = vec![];
      let  _twiter = Tuit {
             nombre:    Some(_req.nombre.to_string().into()),
             comentario: Some(_req.comentario.to_string().into()),
             fecha:   NaiveDate::parse_from_str(&_req.fecha, "%d/%m/%Y").unwrap(),
             hashtags:Some(concatenarcadenas(_req.hashtags.to_vec()).to_string().into()) ,
             upvotes: _req.upvotes,
             downvotes: _req.downvotes
                                       };
        insert_data.push(_twiter);
        conn.exec_batch(
           r"INSERT INTO OLIMPIC (nombre, comentario, fecha,hashtags,upvotes,downvotes)
           VALUES (:nombre, :comentario, :fecha,:hashtags,:upvotes,:downvotes)",
           insert_data.iter().map(|p| params! {
                  "nombre" => &p.nombre,
                  "comentario" => &p.comentario,
                  "fecha" => &p.fecha,
                  "hashtags" => &p.hashtags,
                  "upvotes" => p.upvotes,
                  "downvotes" => p.downvotes,})).unwrap();
                  println!("Se inicio La carga en mysql");
                  CONTADORSQLDB+=1;
                  let start = Instant::now();
                  expensive_function();
                  let duration = start.elapsed();
                  SEGUNDOSMYSQL  = duration.as_secs();

                }
                
                


      println!("Se inicio La carga de Archivos");
  }

  fn expensive_function() {
    thread::sleep(Duration::from_secs(60));
}

  pub async fn finalizar_carga()-> Json<Mensaje> {
   unsafe{
        let mut arreglo:Vec<Notificacion> = Vec::new();
        let mongo_db = Notificacion{
           guardados: CONTADORCOSMODB,
           api: "rust".to_string(),
           tiempo: SEGUNDOSMONGO,
           db: "Cosmodb".to_string(),  
       };
        arreglo.push(mongo_db);
        publicar_mensaje(arreglo);
   }
   /*
   unsafe{
    let mut vector:Vec<Notificacion> = Vec::new();
    let mysqldb = Notificacion{
       guardados:  CONTADORSQLDB,
       api: "rust".to_string(),
       tiempo:   SEGUNDOSMYSQL ,
       db: "mysql".to_string(),  
    };
    vector.push(mysqldb);
    publicar_mensaje(vector);
   }
   */ 

    println!("se Finalizao carga");
    let smsjson = Mensaje { mensaje: "se Finalizo carga!".to_string() };
    return Json(smsjson);  
}







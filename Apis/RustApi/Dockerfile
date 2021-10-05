FROM rust:1.55.0
WORKDIR /usr/src/myapp
COPY . .
RUN cargo install --path .
#RUN cargo run .
#RUN cargo build --release
CMD ["cargo", "run"]
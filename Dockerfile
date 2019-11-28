############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git



WORKDIR $GOPATH/src/pruebadocker/
COPY . .

# Fetch dependencies.
# Using go get.
# go mod sería posible pero no me ha salido fácil en las pruebas, vuelvo a go get
RUN go get -v github.com/gorilla/mux

# RUN go install -v goinsta.v2/examples/show-latest-image/main.go
# RUN go build -o /go/bin/main goinsta.v2/examples/show-latest-image/main.go 
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/main goinsta.v2/examples/show-latest-image/main.go
# COPY /main /go/bin/main
#Se comprueba que el binario está donde debería estar
RUN ls -la /go/bin 

############################
# STEP 2 build a small image
############################
FROM scratch

WORKDIR $GOPATH/src/pruebadocker/
# Copy our static executable. (AHORA SÍ ES STATIC ¬¬)
COPY --from=builder /go/bin/main /go/bin/prueba

# CMD [ "/go/bin/prueba" ]
ENTRYPOINT ["/go/bin/prueba", "--port", "5000"]
# Ayuda de: https://forums.docker.com/t/how-to-expose-port-on-running-container/3252/6
# Aquí aprendí a crear imágenes light de go para Docker y a compilar el ejecutable sin linkeo dinámico:
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
# Esta pregunta de stackoverflow solucionó mis problemas: https://stackoverflow.com/questions/56832363/docker-standard-init-linux-go211-exec-user-process-caused-no-such-file-or-di
# Resulta que el binario que estaba creando no era de linkeo estático y por eso me daba este error:
# standard_init_linux.go:211: exec user process caused "no such file or directory"
# flags de compilación: https://stackoverflow.com/questions/22267189/what-does-the-w-flag-mean-when-passed-in-via-the-ldflags-option-to-the-go-comman



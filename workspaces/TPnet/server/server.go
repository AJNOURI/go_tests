package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type server struct {
	logger     *log.Logger
	listener   net.Listener
	register   chan net.Conn
	unregister chan net.Conn
	broadcast  chan string
}

// ListenAndServe démarre le serveur et bloque durant toute son exécution.
func ListenAndServe(addr string) error {
	s := &server{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		// FIXME instancier les channels nécessaires au fonctionnement du server
	}

	// FIXME utiliser le package net pour créer un nouveau Listener
	//       le Listener doit écouter en TCP sur l'adresse addr

	s.listener = listener

	_, port, err := net.SplitHostPort(s.listener.Addr().String())
	if err != nil {
		s.logger.Printf("Lecture du port échouée: %s", err)
	}
	s.logger.Printf("En attente de connexions sur le port %s", port)

	go s.acceptConnections()

	s.serve()

	return nil
}

func (s *server) serve() {
	var clients = make(map[net.Conn]struct{})

	for {
		select {
		case c := <-s.register:
			clients[c] = struct{}{}
			s.logger.Println("Nouvelle connexion")

		case message := <-s.broadcast:
			// FIXME diffuser le message à tous les clients connectés

		case conn := <-s.unregister:
			delete(clients, conn)
			s.logger.Println("Connexion fermée")

		}
	}
}

func (s *server) acceptConnections() {
	for {
		// FIXME utiliser le Listener du server pour accepter les nouvelles connexions

		go s.manageConnection(conn)
	}
}

func (s *server) manageConnection(conn net.Conn) {
	if _, err := fmt.Fprintln(conn, "Bienvenue ! Merci de saisir votre nom :"); err != nil {
		s.logger.Printf("Envoi de message échoué: %s", err)
		return
	}

	// FIXME utiliser le package bufio pour créer un Reader à partir de la connexion
	//       le Reader permettra de lire le texte envoyé par le client

	username, err := reader.ReadString('\n')
	if err != nil {
		s.logger.Println(fmt.Sprintf("Lecture du nom échouée: %s", err))
		return
	}
	username = username[:len(username)-1]

	s.broadcast <- fmt.Sprintf("%s s'est connecté", username)
	// FIXME ajouter la connexion à la liste des clients

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			// FIXME retirer la connexion de la liste des clients
			s.broadcast <- fmt.Sprintf("%s s'est déconnecté", username)
			break
		}
		if err != nil {
			s.logger.Printf("Réception du message échouée: %s", err)
		}
		message = message[:len(message)-1]

		// FIXME diffuser le message de l'utilisateur à tous les clients
	}
}

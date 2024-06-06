package utils

import (
    "gopkg.in/mail.v2"
    "log"
)

func EnvoyerEmail(destinataire, sujet, corps string) error {
    m := mail.NewMessage()
    m.SetHeader("From", "testrendu1@example.com")
    m.SetHeader("To", destinataire)
    m.SetHeader("Subject", sujet)
    m.SetBody("text/plain", corps)

    d := mail.NewDialer("localhost", 587, "user1@example.com", "password1")

    if err := d.DialAndSend(m); err != nil {
        log.Println("Erreur lors de l'envoi de l'email:", err)
        return err
    }
    return nil
}

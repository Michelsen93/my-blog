# Cloud for en nybegynner

Et nytt kapittel med for meg er å tenke på maskinvaren som faktisk kjører koden jeg skriver. Jeg kan ikke så mye om dette, så det er fint å ha tjenester som leverer en enklere inngang til å ta i bruk servere og alt knyttet til det som Azure i dette tilfellet. Hvorfor Azure? Inntrykket mitt er at det er det som blir brukt stort sett i utviklermiljøet jeg er eksponert til, samt at jeg har en visual studio lisens som gir meg et generøst budsjett som jeg kan benytte til å leke meg med. 

Jeg har til nå satt opp en web app som peker på en docker container som eksisterer i noe som heter container instances. Dette er en under en ressursgruppe. Det har et eget custom domain med proof of ownership hos domeneshop. For en ganske enkel side, begynner det å bli ganske mange produkter som er i bruk. Jeg har klikket meg igjennom dette bare for å få det opp og kjøre, og det er også dette Microsoft sin dokumentasjon oppfordrer til.

### Click Ops

Å sette opp en cloud ved hjelp av brukergrensesnittet til cloud leverandøren.
Det kan være greit å gjøre i starten for å lære seg hvordan cloud fungerer og få en grunnleggende måte, men det skal ikke mye til før det er så mange tjenester og produkter som eksisterer i miljøet at det ikke kan håndteres manuelt. I tillegg så blir staten til clouden din mest sansynlig udokumentert. 
##### Risiko
Hva skjer hvis noen sletter noe med et uhell? Klarer du å gjenskape det? Hva hvis du trenger en ny tjeneste som er ganske lik en du har? Klarer du å gjenskape det? 

### IaC

Løsningen til dette problemet skal være IaC (Infrastructure as Code). Det vil si at du definerer cloud gjennom kode, og kan oppdatere en cloud på samme måte som når du pusher kode til et repo. Her er det mange alternativer å velge mellom. Microsoft ønsker at jeg bryker arm-templates eller bicep. Men disse produktene er ekslusive for Azure, og jeg føler det er ganske meningsløst å lære seg noe som kun virker i Azure. 
Jeg har derfor gjenskap min cloud med Terraform, noe som skal fungere med flere leverandører. (Men i det jeg skriver dette, ser jeg at de ikke har den beste lisensen for bruk av bedrifter, så jeg vil nok i fremtiden velge noe annet som OpenTofu eller Pulumi). Terraform har fungert fint, og gir meg mye bedre selvtillit på at min cloud er implementert på i det minste en helt OK måte.

### Take away

Bruk click ops så lite som mulig og ta i bruk IaC. Det er bedre og lettere.


Takk for tiden
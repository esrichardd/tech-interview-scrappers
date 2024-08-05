# Tech Interview Scrappers

`tech-interview-scrappers` es un servicio desarrollado en Golang que proporciona funcionalidades para obtener información sobre ligas y partidos de fútbol utilizando un servicio de scrapeo al sitio LiveScore. Este servicio corre en el puerto 3003.

## Información General

- **Puerto**: 3003

## Operaciones Disponibles

### Inicializar Datos

Para inicializar los datos en el sistema, puedes utilizar el siguiente comando `curl`:

```bash
curl --location --request POST 'http://localhost:3003/initialize'
```

### Obtener Información y Equipos de la Liga Italiana

Para obtener información general y los equipos de la liga italiana, utiliza el siguiente comando `curl`:

```bash
curl --location 'http://localhost:3003/league'
```

### Obtener Partidos de la Liga Italiana

Para obtener los partidos de la liga italiana, utiliza el siguiente comando `curl`:

```bash
curl --location 'http://localhost:3003/league/matches'
```

### Obtener Partidos de un Equipo

Para obtener los partidos de un equipo, utiliza el siguiente comando `curl`. Reemplaza `:externalTeamId` con el ID externo del equipo que deseas consultar:

```bash
curl --location 'http://localhost:3003/team/:externalTeamId/matches' \
--header 'Content-Type: application/json'
```

### Obtener Todas las Incidencias de un Partido

Para obtener todas las incidencias de un partido, utiliza el siguiente comando `curl`. Reemplaza `:externalMatchId` con el ID externo del partido que deseas consultar:

```bash
curl --location 'http://localhost:3003/match/:externalMatchId/scorebard' \
--header 'Content-Type: application/json'
```

## Contacto

Para cualquier pregunta o problema, puedes contactarme en espinozar1994@gmail.com

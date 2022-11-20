if ([string]::IsNullOrEmpty($Env:POSTGRES_PASSWORD)) {
    echo "Error: environment variable POSTGRES_PASSWORD is not set"
    exit
}

if (((-not (Test-Path generator/out/bandits.csv)) -or (-not (Test-Path generator/out/factions.csv))) -or (-not (Test-Path generator/out/influence.csv))) {
    echo "Error: generator output is not valid (searching ./generator/out/*.csv)"
    exit
}

docker-compose -f compose.yml up --detach
if (-not $?) {
    echo "Error: failed to launch docker container"
    exit
}

docker exec bmstu_db_postgres_1 /bin/sh -c 'psql -U jora -d bandits </var/lib/postgresql/data/sql/initialize.sql'

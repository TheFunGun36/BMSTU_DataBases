create table public.influence(
	"id" int primary key generated always as identity,
	"influence" text
);

create table public.bandits(
	"id" int primary key generated always as identity,
	"nickname" text,
	"lastname" text,
	"firstname" text,
	"midname" text,
	"birth" date,
	"gender" bool,
	"influence" int
		references public.influence("id")
);

create table public.factions(
	"id" int primary key generated always as identity,
	"name" text,
	"founded" int,
	"leader" int
		references public.bandits("id"),
	"location" text,
	"influence" int
		references public.influence("id")
);

alter table public.bandits
	add column "faction" int 
		references public.factions("id");

copy influence
	from '/var/lib/postgresql/data/generator/out/influence.csv'
	delimiter ',' csv header;

copy bandits("nickname","lastname","firstname","midname","birth","gender","influence")
	from '/var/lib/postgresql/data/generator/out/leaders.csv'
	delimiter ',' csv header
	null as 'null';

copy bandits("nickname","lastname","firstname","midname","birth","gender","influence")
	from '/var/lib/postgresql/data/generator/out/bandits.csv'
	delimiter ',' csv header;

copy factions("name","founded","leader","location","influence")
	from '/var/lib/postgresql/data/generator/out/factions.csv'
	delimiter ',' csv header
	null as 'null';

update bandits
	set "faction"= 1 + floor(random()*(select count(*) from public.factions))
	where "faction" is null;

update bandits b
	set "faction" = (select id from factions where b.id = leader)
	where b.id in (select leader from factions);
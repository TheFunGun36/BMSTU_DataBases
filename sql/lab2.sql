
-- 1. Инструкция SELECT, использующая предикат сравнения
select nickname, birth from bandits
where birth < date('1955-01-01');
	

-- 2. Инструкция SELECT, использующая предикат BETWEEN
select nickname, birth from bandits
where birth between date('1975-01-01') and date('1980-01-01');
	
-- 3. Инструкция SELECT, использующая предикат LIKE.
-- ищем всех Мухиных, вне зависимости от пола
select nickname, lastname from bandits
where lastname like 'Мухин%';
	
-- 4. Инструкция SELECT, использующая предикат IN с вложенным подзапросом.
-- Прозвища всех лидеров
select nickname from bandits
where id in (
	select leader from factions)
	
-- 5. Инструкция SELECT, использующая предикат EXISTS с вложенным подзапросом.
-- Прозвища всех лидеров
select nickname from bandits b
where exists (
	select id from factions
	where leader = b.id)
	
-- 6. Инструкция SELECT, использующая предикат сравнения с квантором
-- Все группировки, в которых нет ни одного члена, профиль которого совпадает с профилем группировки
select * from factions f
where influence != all (
	select influence from bandits
	where bandits.faction = f.id);
		
-- 7. Инструкция SELECT, использующая агрегатные функции в выражениях столбцов
-- Дата рождения старшего члена Монолита
select min(birth) from bandits
where faction = (
	select id from factions where name = 'Монолит');
	
-- 8. Инструкция SELECT, использующая скалярные подзапросы в выражениях столбцов
-- Возраст старшего члена для каждой группировки
select id, name,
	(select min(birth) from bandits
		where faction = f.id) as oldest
from factions f;
	
-- 9. Инструкция SELECT, использующая простое выражение CASE
select nickname as "Прозвище",
	case gender
		when true then 'Женский'
		when false then 'Мужской'
		else 'Неизвестно'
	end as "Пол"
from bandits;
	
-- 10. Инструкция SELECT, использующая поисковое выражение CASE
select nickname as "Прозвище", birth as "Родился",
	case
		when birth < date('1960-1-1') and gender then 'Застала события Yakuza 0'
		when birth < date('1960-1-1') then 'Застал события Yakuza 0'
		when birth < date('1970-1-1') and gender then 'нууу, тут врятли есть пороховницы'
		when birth < date('1970-1-1') then 'ЕСТЬ ЕЩЁ ПОРОХ В ПОРОХОВНИЦАХ'
		when birth < date('1980-1-1') then 'эмм'
		when birth > date('2000-1-1') then 'Чёт молодой'
		else 'фантазия END'
	end as "Статус"
from bandits;
	
-- 11. Создание новой временной локальной таблицы из результирующего набора данных инструкции SELECT.
create temp table leaders as
	select *
	from bandits b
	where exists (
		select id from factions
		where leader = b.id);

-- 12. Инструкция SELECT, использующая вложенные коррелированные подзапросы в качестве производных таблиц в предложении FROM
-- Прозвища всех лидеров
select nickname from bandits b
where exists (
	select id from factions
	where leader = b.id)
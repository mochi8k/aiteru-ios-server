create table places(
    `id` int(11) unsigned not null auto_increment,
    `place_name` varchar(255) not null,
    `created_at` datetime,
    `created_by` int(11),
    `updated_at` datetime,
    `updated_by` int(11),
    primary key(id)
);

create table place_owners(
    `place_id` int(11) unsigned not null,
    `owner_id` int(11) unsigned not null,
    primary key(place_id, owner_id)
);

create table place_collaborators(
    `place_id` int(11) unsigned not null,
    `collaborator_id` int(11) unsigned not null,
    primary key(place_id, collaborator_id)
);

create table place_status(
    `place_id` int(11) unsigned not null,
    `is_open` boolean not null,
    `updated_at` datetime,
    `updated_by` int(11)
);


select
p.id,
p.place_name,
group_concat(distinct u1.user_name) as owner_names,
group_concat(distinct u2.user_name) as collaborator_names,
u3.user_name as created_by,
p.created_at,
u4.user_name as updated_by,
p.updated_at
from places as p
inner join place_owners as po on p.id = po.place_id
inner join users as u1 on u1.id = po.owner_id
inner join place_collaborators as pc on p.id = pc.place_id
inner join users as u2 on u2.id = pc.collaborator_id
inner join users as u3 on p.created_by=u3.id
left outer join users as u4 on p.updated_by=u3.id
group by p.id, u3.user_name, u4.user_name
;

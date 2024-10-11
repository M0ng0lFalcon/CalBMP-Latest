## new version to get basic data

```sql
-- get county names
select distinct "NAME"
from zip_county
order by "NAME";

-- get zipcode of county name
select distinct "ZIP_CODE"
from zip_county
where "NAME" = 'Monterey'
  and "ZIP_CODE" in (
    select zip_code
    from zip_mukey_cokey
    where zip_code = zip_county."ZIP_CODE"
      and compname is not null
);

-- get soil name(compname) by zipcode
select distinct compname
from zip_mukey_cokey
where zip_code = '93204'
  and compkind = 'Series'
```
from datetime import date, timedelta
import datetime
import random

COMPANIES_N = 1000
SHAREHOLDERS_N = 3000
AGREEMENTS_N = 4000


def read_data(filename):
    f = open(filename, 'r')
    data = f.read().split('\n')
    f.close()
    return data[:-1]


def dates(start, end):
    res = []
    res.append(start)
    while start != end:
        start += timedelta(days=1)
        res.append(start)
    return res

def create_shareholders(n, filename):
    bdays = []

    male_names = read_data("mnames.txt")
    female_names = read_data("fnames.txt")
    male_lastnames = read_data("mlastnames.txt")
    female_lastnames = read_data("flastnames.txt")
    male_father_names = read_data("mfathernames.txt")
    female_father_names = read_data("ffathernames.txt")

    passports = random.sample(range(1000000000, 10000000000), n)

    sex = ["м", "ж"]

    emails = read_data("emails.txt")
    random.shuffle(emails)
    domains = read_data("domains.txt")

    phone_numbers = random.sample(range(1000000000, 10000000000), n)

    bdates = dates(date(1960, 1, 1), date(2000, 12, 31))

    cities = read_data("cities.txt")

    f = open(filename, 'w')
    for i in range(n):
        s = random.choice(sex)
        if (s == 'ж'):
            name = random.choice(female_lastnames) + " " + \
                random.choice(female_names) + " " + random.choice(female_father_names)
        else:
            name = random.choice(male_lastnames) + " " + \
                random.choice(male_names) + " " + random.choice(male_father_names)
        passport = str(passports[i]) 

        bday = random.choice(bdates)
        bdays.append(date(bday.year + 18, bday.month, (bday.day - 1 if (bday.month == 2 and bday.day == 29) else bday.day)))
        bday = (str(bday))

        city = random.choice(cities)
        phone_number = str(phone_numbers[i]) 
        email = emails[i] + random.choice(domains)

        f.write(str(i + 1) + ";" +
                name + ";" + 
                passport + ";" + 
                s + ";" +
                bday + ";" +
                city + ";" +
                phone_number + ";" +
                email + "\n")
    f.close()

    return bdays


def create_companies(n, filename):
    regdays = []

    male_names = read_data("mnames.txt")
    female_names = read_data("fnames.txt")
    male_lastnames = read_data("mlastnames.txt")
    female_lastnames = read_data("flastnames.txt")
    male_father_names = read_data("mfathernames.txt")
    female_father_names = read_data("ffathernames.txt")
    sex = ["м", "ж"]

    companies_names = read_data("cnames.txt")
    random.shuffle(companies_names)
    
    psrns = random.sample(range(1000000000000, 10000000000000), n)

    rdates = dates(date(1995, 1, 1), date(2014, 12, 31))

    cities = read_data("cities.txt")

    f = open(filename, 'w')
    for i in range(n):
        s = random.choice(sex)
        if (s == 'ж'):
            name = random.choice(female_lastnames) + " " + \
                random.choice(female_names) + " " + random.choice(female_father_names)
        else:
            name = random.choice(male_lastnames) + " " + \
                random.choice(male_names) + " " + random.choice(male_father_names)

        company_name = companies_names[i]
        psrn = str(psrns[i]) 

        regday = random.choice(rdates)
        regdays.append(regday)
        regday = (str(regday))#[:10]

        city = random.choice(cities)
    
        f.write(str(i + 1) + ";" +
                company_name + ";" + 
                psrn + ";" + 
                name + ";" +
                city + ";" +
                regday + "\n")
    f.close()

    return regdays

def create_objects(n, filename):
    objsubtypes = ["кв", "мм", "кп", ""]
    objtypes = ["ж", "н"]

    appartments_min = 30
    car_place_min = 13
    storage_room_min = 2

    appartments_max = 150
    car_place_max = 22
    storage_room_max = 7

    cities = read_data("cities.txt")
    streets = read_data("streets.txt")
    buildings = read_data("buildings.txt")

    maxnum = 100

    f = open(filename, 'w')
    for i in range(n):
        objtype = random.choice(objtypes)

        if (objtype == "ж"):
            objsubtype = random.choice(["кв", ""])
            area = str("%.1f" % random.uniform(appartments_min, appartments_max))
        else:
            objsubtype = random.choice(["мм", "кп", ""])
            if (objsubtype == "мм"):
                area = "%.1f" % random.uniform(car_place_min, car_place_max)
            elif (objsubtype == "кп"):
                area = "%.1f" % random.uniform(storage_room_min, storage_room_max)
            else:
                area = "%.1f" % random.uniform(storage_room_min, appartments_max)

        city = random.choice(cities)
        
        street = random.choice(streets)
        house = str(random.randint(1, maxnum))
        objnum = str(random.randint(1, maxnum))
        building = random.choice([random.choice(buildings), ""])
    
        f.write(str(i + 1) + ";" +
                objtype + ";" + 
                objsubtype + ";" + 
                area + ";" +
                city + ";" +
                street + ";" +
                house + ";" +
                building + ";" +
                objnum + "\n")
    f.close()

def create_agreements(n, filename, rdays, bdays):
    letters = "ЙЦУКЕНГШЩЗХЭЖДЛОРПАВЫФЯЧСМИТБЮ"

    obj_ind = []
    for i in range(n):
        obj_ind.append(i + 1)

    agr_numbers = random.sample(range(10000, 99999), n)

    a = []
    b = []
    k = 0

    f = open(filename, 'w')
    for i in range(n):
        agr_number = random.choice(letters) + random.choice(letters) + "-" + str(agr_numbers[i])
        price = "%.2f" %random.uniform(500000, 50000000)

        obj_id = random.choice(obj_ind)
        obj_ind.remove(obj_id)
        obj_id = str(obj_id)

        c_id = random.randint(1, COMPANIES_N)
        s_id = random.randint(1, SHAREHOLDERS_N)

        start_date = max(rdays[c_id - 1], bdays[s_id - 1])
        conc_date = random.choice(dates(start_date, date.today()))

        deadline_date = random.choice(dates(conc_date, date(conc_date.year + 4, conc_date.month, round(conc_date.day / 2) + 1)))


        end_date = min(date(deadline_date.year + 3, deadline_date.month, (deadline_date.day - 1 if (deadline_date.month == 2 and deadline_date.day == 29) else deadline_date.day)), date.today())
        rec_date = random.choice(dates(conc_date, end_date))

        j_id = ""
        if (rec_date > deadline_date):
            a.append(i)
            b.append(deadline_date)
            j_id = k + 1
            k += 1
            
        if (rec_date <= date.today() and deadline_date.year >= 2019):
            rec_date = random.choice([str(rec_date), ""])
        else:
            rec_date = str(rec_date)

        conc_date = str(conc_date)
        deadline_date = str(deadline_date)

        f.write(str(i + 1) + ";" +
                str(c_id) + ";" + 
                str(s_id) + ";" + 
                str(obj_id) + ";" +
                str(j_id) + ";" +
                agr_number + ";" +
                price + ";" +
                conc_date + ";" +
                deadline_date + ";" +
                rec_date + "\n")
    f.close()

    return a, b, k
    
def create_judgement(n, filename, a, b):
    male_names = read_data("mnames.txt")
    female_names = read_data("fnames.txt")
    male_lastnames = read_data("mlastnames.txt")
    female_lastnames = read_data("flastnames.txt")
    male_father_names = read_data("mfathernames.txt")
    female_father_names = read_data("ffathernames.txt")

    sex = ["м", "ж"]

    f = open(filename, 'w')
    for i in range(n):
        s = random.choice(sex)
        if (s == 'ж'):
            name = random.choice(female_lastnames) + " " + \
                random.choice(female_names) + " " + random.choice(female_father_names)
        else:
            name = random.choice(male_lastnames) + " " + \
                random.choice(male_names) + " " + random.choice(male_father_names)

        session = random.choice(dates(b[i], date(b[i].year + 2, b[i].month,round(b[i].day / 2) + 1)))

        claim = random.uniform(500000, 1000000)

        if (session > date.today()):
            decision = ""
        else:
            decision = "%.2f" %(random.randint(0, 150) * claim / 100)
            
        claim = "%.2f" %claim

        session = (str(session))


        f.write(str(i + 1) + ";" +
                name + ";" + 
                session + ";" +
                claim + ";" +
                decision + "\n")

    f.close()

def main():
    b = create_shareholders(SHAREHOLDERS_N, "shareholders.csv")
    r = create_companies(COMPANIES_N, "companies.csv")
    create_objects(AGREEMENTS_N, "objects.csv")
    c, d, e = create_agreements(AGREEMENTS_N, "agreements.csv", r, b)
    create_judgement(e, "judgement.csv", c, d)
    

if __name__ == "__main__":
    main()

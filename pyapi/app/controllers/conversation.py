from app.models import robot

def generate_robot():
    diagnosis_robot = robot.DiagnosisRobot()
    return diagnosis_robot

def only_hello():
    diagnosis_robot = robot.DiagnosisRobot()
    return diagnosis_robot.say_only_hello()

def talk_about_diagnosis(r, data):
    a = r.hello(data)
    b = r.ask_only_feeling(data)
    return b

def talk_about_dignosis_second(r, choise):
    r.ask_user_feeling(choise)
    c = r.ask_only_sleepy()
    return c

def talk_about_dignosis_third(r, choise):
    r.ask_user_sleep(choise)
    d = r.ask_only_worried()
    return d

def talk_about_dignosis_forth(r, choise):
    r.ask_user_worry(choise)
    d = r.generate_diagnosis_result()
    return d

def talk_about_dignosis_final(r):
    d = r.thank_you()
    return d
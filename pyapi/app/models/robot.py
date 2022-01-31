from app.models import diagnosis
from app.views import console


DEFAULT_ROBOT_NAME = 'DiagnosisRobot'

class Robot(object):

    def __init__(self, name=DEFAULT_ROBOT_NAME, user_name='hoge',
                 speak_color='green') -> None:
        self.name = name
        self.user_name = user_name
        self.speak_color = speak_color
    
    def say_only_hello(self):
        temp_path = console.find_template('hello.txt')
        with open(temp_path, 'r') as file:
            data = file.read().replace('\n', '')
            robot_name = DEFAULT_ROBOT_NAME
            data = 'Hello, I am %s. What is your name?' % robot_name
            return data

    def hello(self, u_name):
        while True:
            user_name = u_name
            if user_name:
                self.user_name = user_name.title()
                temp_path = console.find_template('hello.txt')
                with open(temp_path, 'r') as file:
                    data = file.read().replace('\n', '')
                    return data
            

class DiagnosisRobot(Robot):

    def __init__(self, name=DEFAULT_ROBOT_NAME) -> None:
        super().__init__(name)
        self.diagnosis_model = diagnosis.DiagnosisModel()
        self.advice = ""
        self.score = 0
    
    def _hello_decorator(func):
        def wrapper(self, hoge=None):
            if not self.user_name:
                self.hello()
            return func(self, hoge)
        return wrapper
    
    def ask_only_feeling(self, msg):
        diagnosis_question = 'Hello, %s. How is your feeling now?\n Please choose number [1: so bad] [2: so so] [3: so good]' % msg
        return diagnosis_question

    @_hello_decorator
    def ask_user_feeling(self, chosed_num):
        diagnosis_question = "are you feeling good?"
        self.diagnosis_model.increment(diagnosis_question, chosed_num)
                
    def ask_only_sleepy(self):
        diagnosis_question = 'Are you sleepy now?\n Please choose number [1: so sleepy] [2: so so] [3: not at all]'
        return diagnosis_question 

    @_hello_decorator 
    def ask_user_sleep(self, chosed_num):
        diagnosis_question = "Are you sleepy?"
        self.diagnosis_model.increment(diagnosis_question, chosed_num)

    def ask_only_worried(self):
        diagnosis_question = 'Do you have something worry?\n Please choose number [1: a lot] [2: so so] [3: not at all]'
        return diagnosis_question 

    @_hello_decorator
    def ask_user_worry(self, chosed_num):
        diagnosis_question = "Do you have some worried?"
        self.diagnosis_model.increment(diagnosis_question, chosed_num)

    def generate_diagnosis_result(self):
        diagnosis_question = 'I generated result! Can I show?[y/n]'
        return diagnosis_question

    @_hello_decorator
    def thank_you(self, hoge):
        print(hoge)
        result = ""
        score = 0
        previous_data = self.diagnosis_model.load_data()
        for _, v in previous_data.items():
            try:
                int(v)
            except:
                return "error: value is not intger"
            score += v
        if score in [8, 9]:
            result = "you are perfect condition today! trade!"
        elif score in [3, 4]:
            result = "you are so bad condition ,,, stop trading today"
        else:
            result = "you are so so so trade carefully"
        diagnosis_question = 'Result %d/9 -> %s good bye! (fill [end] and enter)' %(score, result)
        return diagnosis_question 
    
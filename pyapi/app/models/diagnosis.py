import pathlib
import os
import collections
import csv

GIVEN_QUESTION_COLUMN = 'QUESTION'
USER_CHOISE_COLUMN = 'ANSWER'
DIAGNOSIS_RUSULT = 'diagnosis.csv'

class CsvModel(object):

    def __init__(self, csv_file) -> None:
        self.csv_file = csv_file
        if not os.path.exists(csv_file):
            pathlib.Path(csv_file).touch()

class DiagnosisModel(CsvModel):

    def __init__(self, csv_file=None, *args, **kwargs) -> None:
        if not csv_file:
            csv_file = self.get_csv_file_path()
        super().__init__(csv_file, *args, **kwargs)
        self.column = [GIVEN_QUESTION_COLUMN, USER_CHOISE_COLUMN]
        self.data = collections.defaultdict(int)
    
    def get_csv_file_path(self):
        csv_file_path = None

        try:
            import settings
            if settings.CSV_FILE_PATH:
                csv_file_path = settings.CSV_FILE_PATH
        except ImportError:
            pass

        if not csv_file_path:
            csv_file_path = DIAGNOSIS_RUSULT
        return csv_file_path
    
    def load_data(self):
        with open(self.csv_file, 'r+') as csv_file:
            reader = csv.DictReader(csv_file)
            for row in reader:
                self.data[row[GIVEN_QUESTION_COLUMN]] = int(
                    row[USER_CHOISE_COLUMN])
            return self.data

    def save(self):
        with open(self.csv_file, 'w+') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=self.column)
            writer.writeheader()

            for name, count in self.data.items():
                writer.writerow({
                    GIVEN_QUESTION_COLUMN: name,
                    USER_CHOISE_COLUMN: count
                })
    
    def increment(self, name, num):
        previous_data = self.load_data()
        for k, v in previous_data.items():
            self.data[k] = v
        self.data[name.title()] = num
        print(self.data)
        self.save()

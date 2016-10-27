import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
import testutil as tu

def main():
	failcount = 0
	d = tu.newtest()
	
	tu.endtest(d)
	sys.exit(failcount)

if __name__ == '__main__':
	main()


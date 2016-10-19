import sys
import os
import time
import subprocess as sp
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
import testutil as tu
import subprocess as sp

_service_pid = 0

def servTest(d):
	try:
		wait = WebDriverWait(d, timeout=5, poll_frequency=0.2)
		inp = d.find_element_by_xpath("//div[@class='input-group']/input")
		inp.send_keys("You are a winner you have been specially selected to win 1000 cash or 2000 reward. Speak to a live operator to claim call 087123002209am to 7pm cost 10p")
		d.find_element_by_xpath("//input[@id='predict-btn']").click()
		wait.until(lambda x: x.find_element_by_class_name("labelHighlight"))
		x = d.find_element_by_class_name("labelHighlight")
		if x.text != "spam":
			print "failed to classify spam correctly"
			return False
		return True
	except:
		print "Failed to navigate prediction page"
		return False


def setup():
	gopath = os.environ['GOPATH']
	path = gopath + '/src/github.com/h2oai/steam/prediction-service-builder/examples'
	os.chdir(gopath + '/src/github.com/h2oai/steam/prediction-service-builder/examples/spam-detection-python')
	ret = sp.check_output('sh example-python.sh', shell=True)
	proc = sp.Popen("/usr/bin/env java -jar ../jetty-runner-8.1.14.v20131031.jar --port 55001 example-python.war > /dev/null 2>&1 &", shell=True)
	time.sleep(5)
	_service_pid = proc.pid + 1
	driver = tu.newtest()
	driver.get("http://localhost:55001/")
	return driver

def main():
	failcount = 0
	d = setup()

	if not servTest(d):
		failcount = failcount + 1
	
	tu.endtest(d)
	sys.exit(failcount)
	

if __name__ == '__main__':
	main()


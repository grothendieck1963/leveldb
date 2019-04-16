base = "test-stuff copy"

import os
from subprocess import Popen, PIPE
import json
import ast
for f in os.listdir(base):
  if f.endswith(".ldb"):
    process = Popen(["ldbdump", os.path.join(base, f)], stdout=PIPE)
    (output, err) = process.communicate()
    exit_code = process.wait()
    for line in (output.split("\n")[1:]):
      if line.strip() == "": continue
      parsed = ast.literal_eval("{" + line + "}")
      key = parsed.keys()[0]
      print json.dumps({ "key": key.encode('string-escape'), "value": parsed[key] })

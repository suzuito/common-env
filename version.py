import re
import sys

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage version.py <Filename for output>")
        exit(1)
    outFilename = sys.argv[1]
    latest_version = ''
    latest_messages = []
    with open('RELEASE.md') as f:
        for l in f.readlines():
            l = l.rstrip()
            g = re.search('^# (.+)$', l)
            if g is None:
                if latest_version != '':
                    latest_messages.append(l)
            else:
                if latest_version == '':
                    latest_version = g.group(1)
                else:
                    break
    print(latest_version)
    print(latest_messages)
    with open('{}'.format(outFilename), 'w') as f:
        f.write(latest_version)
    with open('{}.md'.format(outFilename), 'w') as f:
        f.write('# ' + latest_version)
        f.write('\n')
        for l in latest_messages:
            f.write(l)
            f.write('\n')


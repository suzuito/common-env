import re

if __name__ == "__main__":
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
    with open('LATEST_VERSION', 'w') as f:
        f.write(latest_version)
    with open('LATEST_VERSION.md', 'w') as f:
        f.write('# ' + latest_version)
        f.write('\n')
        for l in latest_messages:
            f.write(l)
            f.write('\n')
